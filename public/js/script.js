$('.logout').click(function () {
  $('#logout-form').submit();
});

$('.show-loading').submit(function () {
  var loading = $('.loading');
  loading.removeClass('dispNone');
});

$('.connect-git-hub').submit(function () {
  new PNotify({
    title: 'Connecting To GitHub',
    text: 'This process needs time a lot.',
    type: 'info',
    hide: false,
    styling: 'bootstrap3'
  });
});