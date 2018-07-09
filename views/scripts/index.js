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
                    // $('<div><img src="assets/dir.png" class="img-icon"><span class="file-name"><a href="javascript:void(0)">'+data[ix].name+'</a></span></div>').appendTo($('body'))
                    $('<div><em class="glyphicon glyphicon-folder-open"></em><span class="file-name"><a href="javascript:void(0)">'+data[ix].name+'</a></span></div>').appendTo($('body'))
                }else{
                    $('<div><em class="glyphicon glyphicon-file"></em><span class="file-name">'+data[ix].name+'</span><span class="file-size">'+data[ix].size+'</span><a class="btn-download" href="#">下载</a></div>').appendTo($('body'))
                }
            }
            $('a').click(function(){
                go($(this).html())
            })
        }
    })
})