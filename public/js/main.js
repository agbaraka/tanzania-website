/* global jQuery */
jQuery(document).ready(function($){
	var menu = $('.menu'),
		foldingPanel = $('.folding-panel'),
		mainContent = $('.main');

	// open folding content
	menu.on('click', 'a', function(event){
		event.preventDefault();
		openItemInfo($(this).attr('href'));
	});

	// close folding content
	foldingPanel.on('click','.close', function(event){
		event.preventDefault();
		toggleContent('', false);
	});
	menu.on('click', function(event){
		//detect click on .menu::before when the .folding-panel is open
		if($(event.target).is('.menu') && $('.fold-is-open').length > 0) toggleContent('',false);
	});

	function openItemInfo(url){
		var mq = viewportSize();
		if(menu.offset().top > $(window).scrollTop() && mq != 'mobile'){
			//if content is visible above the .menu - scroll before opening the folding panel
			$('body,html').animate({
				'scrollTop': menu.offset().top
			}, 100, function(){
				toggleContent(url, true);
			});
		}else if(menu.offset().top + menu.height() < $(window).scrollTop() + $(window).height() && mq != 'mobile'){
			//if content id visible below the .menu - scroll before opening the folding panel
			$('body,html').animate({
				'scrollTop': menu.offset().top + menu.height() - $(window).height()
			}, 100, function(){
				toggleContent(url,true);
			})
		}else{
			toggleContent(url, true);
		}
	}

	function toggleContent(url, bool){
		if(bool){
			//load and show new content
			var foldingContent = foldingPanel.find('.fold-content');
			foldingContent.load(url+" .fold-content > *", function(event){
				setTimeout(function(){
					$('body').addClass('overflow-hidden');
					foldingPanel.addClass('is-open');
					mainContent.addClass('fold-is-open');
				},100);
			})
		}else{
			//close the folding panel
			var mq = viewportSize();
			foldingPanel.removeClass('is-open');
			mainContent.removeClass('fold-is-open');


			(mq == 'mobile' || $('.no-csstransitions').length > 0)
				//according to the mq, immediately remove the .overfow-hidden or wait for the end of the animation
				? $('body').removeClass('overflow-hidden')

				: mainContent.find('.item').eq(0).one('webkitTransitionEnd otransitionend oTransitionEnd msTransitionEnd transitionend', function(){
					$('body').removeClass('overflow-hidden');
					mainContent.find('.item').eq(0).off('webkitTransitionEnd otransitionend oTransitionEnd msTransitionEnd transitionend')
				})

		}
	}

	function viewportSize(){
		// retrive the content value of .main::before to check the actual mq
		return window.getComputedStyle(document.querySelector('.main'), '::before').getPropertyValue('content').replace(/"/g, "").replace(/'/g, "");
	}
})
