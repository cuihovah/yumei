function go(path){
    try{
        var qs = window.location.href.split('#')[0]
        qs = qs.split('?')[1]
        if(!qs){
            qs = ''
        }
        qs = qs.split('&')
        var query = {}
        var pathstr = ''
        for(var ix in qs){
            var kv = qs[ix].split('=')
            if(kv.length > 1){
                query[kv[0]] = kv[1]
            }
        }
        if(query.path){
            pathstr = [query.path, path].join('/')
        }else{
            pathstr = ['/', path].join('')
        }
        pathstr = pathstr.replace(/\/\//g, '/')
        pathstr = ['path', pathstr].join('=')
        window.location.href = [window.location.origin, pathstr].join('?')
    }catch(err){
        console.error(err.message)
    }
}

$(document).ready(function(){
    var query = window.location.href.split('?')[1]
    if(!query){
        query = 'path='
    }
    $.ajax('/files?'+query, {
        'type': 'GET',
        success: function(res){
            var data = res.data
            for(var ix in data){
                if(data[ix].type != 'file'){
                    $('<div><em class="glyphicon glyphicon-folder-open"></em><span class="file-name"><a class="dirname" href="javascript:void(0)">'+data[ix].name+'</a></span></div>').appendTo($('body'))
                }else{
                    $('<div><em class="glyphicon glyphicon-file"></em><span class="file-name">'+data[ix].name+'</span><span class="file-size">'+data[ix].size+'</span><a class="btn-download" href="javascript:void(0)" data="'+data[ix].name+'">下载</a></div>').appendTo($('body'))
                }
            }
            $('.dirname').click(function(){
                go($(this).html())
            })

            $('.btn-download').click(function(){
                var qs = window.location.href.split('#')[0]
                qs = qs.split('?')[1]
                if(!qs){
                    qs = ''
                }
                qs = qs.split('&')
                var query = {}
                var pathstr = ''
                for(var ix in qs){
                    var kv = qs[ix].split('=')
                    if(kv.length > 1){
                        query[kv[0]] = kv[1]
                    }
                }
                var name = $(this).attr('data')
                if(query.path){
                    pathstr = [query.path, name].join('/')
                }else{
                    pathstr = ['/', name].join('')
                }
                pathstr = pathstr.replace(/\/\//g, '/')
                window.location.href = '/files/download?path='+pathstr+'&name='+name
            })
        }
    })
})