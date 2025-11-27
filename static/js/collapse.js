(function ($) {
      const $btn = $('#collapseToggle');
      const $content = $('#collapseContent');
      const $chev = $btn.find('svg.chev');

      // Animation duration in milliseconds
      const DURATION = 300;

      // openCollapse: animate from 0 -> contentHeight, then set height to auto
      function openCollapse() {
        // ensure the element is visible for scrollHeight measurement
        $content.css('display', 'block'); // in case someone set display:none elsewhere
        const fullHeight = $content[0].scrollHeight;

        // start from current height (should be 0), animate to fullHeight
        $content.stop(true).animate({ height: fullHeight }, {
          duration: DURATION,
          easing: 'swing',
          complete: function () {
            // set to auto so it adapts to internal changes after open
            $content.css('height', 'auto');
          }
        });

        // rotate chevron and update aria
        $chev.css('transform', 'rotate(180deg)');
        $btn.attr('aria-expanded', 'true');
      }

      // closeCollapse: animate from current (or auto) height -> 0
      function closeCollapse() {
        // if height is 'auto', set it to actual pixel height before animating
        if ($content.css('height') === 'auto') {
          $content.css('height', $content[0].scrollHeight + 'px');
        }

        // animate to 0
        $content.stop(true).animate({ height: 0 }, {
          duration: DURATION,
          easing: 'swing',
          complete: function () {
            // optional: keep display:block so measurement works later; you can hide if you like
          }
        });

        $chev.css('transform', 'rotate(0deg)');
        $btn.attr('aria-expanded', 'false');
      }

      // Toggle handler
      $btn.on('click', function (e) {
        e.preventDefault();
        const expanded = $btn.attr('aria-expanded') === 'true';
        if (expanded) closeCollapse();
        else openCollapse();
      });

      // keyboard accessibility (Enter/Space)
      $btn.on('keydown', function (e) {
        if (e.key === 'Enter' || e.key === ' ') {
          e.preventDefault();
          $btn.trigger('click');
        }
      });

      // Optional: close when clicking outside
      $(document).on('click', function (e) {
        if (!$(e.target).closest('#collapseContent, #collapseToggle').length) {
          if ($btn.attr('aria-expanded') === 'true') closeCollapse();
        }
      });

      // Example: if content changes dynamically, ensure open state adapts
      // If you programmatically add content and the panel is open, fix height -> auto again:
      const observer = new MutationObserver(() => {
        if ($btn.attr('aria-expanded') === 'true') {
          // short "refresh" so content can expand naturally
          $content.css('height', $content[0].scrollHeight + 'px');
          setTimeout(() => $content.css('height', 'auto'), DURATION + 20);
        }
      });
      observer.observe($content[0], { childList: true, subtree: true, characterData: true });

    })(jQuery);

