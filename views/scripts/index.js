function go(path){
    var query = window.location.href.split('?')[1]
    query = [query, path].join('/')
    window.location.href = [window.location.origin, query].join('?')
}

$(document).ready(function(){
    var query = window.location.href.split('?')[1]
    $.ajax('/files?'+query, {
        'type': 'GET',
        success: function(res){
            var data = res.data
            for(var ix in data){
                if(data[ix].type != 'file'){
                    $('<div><img src="assets/dir.png" style="vertical-align: middle;" height="20px"><a href="javascript:void(0)">'+data[ix].name+'</a></div>').appendTo($('body'))
                }else{
                    $('<div><img src="assets/file.png" style="vertical-align: middle;" height="20px">'+[data[ix].name, data[ix].size, '<a href="#">下载</a>'].join('&nbsp;&nbsp;&nbsp;&nbsp;')+'</div>').appendTo($('body'))
                }
            }
            $('a').click(function(){
                go($(this).html())
            })
        }
    })
})