$(function(){
  var todos = $('[id^=todo-]');

  setInterval(
    function(){
      todos.each(function(index, element){
        var id = $(element).attr('id').replace('todo-', '');

        var now = new Date();
        var createdAt = Date.parse($(`#created-${id}`).text());
        var deadline = Date.parse($(`#deadline-${id}`).text());

        var limit = deadline - now;
        var progress = (limit / (deadline - createdAt)) * 100;

        $(`#timer-${id}`).text(formatLimit(limit));
        $(`#bar-${id}`).css({'width':progress+'%'});
        changeBarColor(`#bar-${id}`, progress);
      });
    }, 1000);
});

function formatLimit(limit){
  const secs = Math.floor(limit / 1000);
  const mins = Math.floor(secs / 60);
  const hours = Math.floor(mins / 60);
  const days = Math.floor(hours / 24);

  // const ms = String(limit % 1000).padStart(3, '0');
  const s = String(secs % 60).padStart(2, '0');
  const m = String(mins % 60).padStart(2, '0');
  const h = String(hours % 24).padStart(2, '0');

  return `${days}d ${h}h ${m}m ${s}s left`;
}

function changeBarColor(progressId, progress){
  if (progress <= 20) {
    $(progressId).removeClass();
    $(progressId).addClass('progress-bar bg-danger');
  }
  else if (progress < 50) {
    $(progressId).removeClass();
    $(progressId).addClass('progress-bar bg-warning');
  }
  else {
    $(progressId).removeClass();
    $(progressId).addClass('progress-bar bg-success');
  }
}
