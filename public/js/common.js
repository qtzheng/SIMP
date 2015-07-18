var request = {
    QueryString: function (val) {
        var uri = window.location.search;
        var re = new RegExp("" + val + "=([^&?]*)", "ig");
        return ((uri.match(re)) ? (uri.match(re)[0].substr(val.length + 1)) : null);
    }
}
var IsOpenWindow = function () {
    var winId = request.QueryString("_winid");
    if (!winId && winId != "") {
        return true;
    }
    else
        return false;
}
function OpenWindow(url, title, iconCls, width, height, data, ondestroy) {
    mini.open({
        url: url,        //页面地址
        title: title,      //标题
        iconCls: iconCls,    //标题图标
        width: width,      //宽度
        height: height,     //高度
        allowResize: true,       //允许尺寸调节
        allowDrag: true,         //允许拖拽位置
        showCloseButton: true,   //显示关闭按钮
        showMaxButton: false,     //显示最大化按钮
        showModal: true,         //显示遮罩
        loadOnRefresh: true,       //true每次刷新都激发onload事件
        onload: function () {       //弹出页面加载完成
            if (typeof data != "undefined") {
                var iframe = this.getIFrameEl();
                //调用弹出页面方法进行初始化
                iframe.contentWindow.InitWindow(data);
            }
        },
        ondestroy: function (data) {  //弹出页面关闭前
            if (typeof ondestroy == "function") {
                data = mini.clone(data);
                ondestroy(data);
            }
        }
    });
}
function CloseOpenWindow(action, data, form) {
    if (form.isChanged) {
        if (action == "close" && form.isChanged()) {
            if (confirm("数据被修改了，是否先保存？")) {
                return false;
            }
        }
    }

    if (window.CloseOwnerWindow && IsOpenWindow())
        return window.CloseOwnerWindow(data);
    //else
    //    window.close();
}
function OpenAddForm(form, win, title, icons, funcBefore) {
    form.clear();
    if (typeof funcBefore == "function") {
        funcBefore();
    }
    win.set({ title: title, iconCls: icons });
    win.showAtPos("center", "middle");
}
function OpenEditForm(url, form, win, title, icons, funcBefore,funcLoad) {
    OpenWaite();
    try {
        form.clear();
        if (typeof funcBefore == "function") {
            funcBefore();
        }
        $.ajax({
            url: url,
            type: "get",
            success: function (text) {
                CloseWaite();
                var data = mini.decode(text);   //反序列化成对象
                form.setData(data);             //设置多个控件数据
                if (typeof funcLoad == "function") {
                    funcLoad(data);
                }
            }, error: function (q, m, e) {
                CloseWaite();
                mini.alert(m);
            }
        });
        win.set({ title: title, iconCls: icons });
        win.showAtPos("center", "middle");
    } catch (e) {
        CloseWaite();
        mini.alert(e.message);
    }
}

function HideWin(win) {
    if (typeof win == "String") {
        mini.get(win).hide();
    }
    else
        win.hide();
}
function OpenWaite(msg) {
    if (!msg) {
        msg = "请等待...";
    }
    mini.mask({
        el: document.body,
        cls: 'mini-mask-loading',
        html: msg
    });
}
function CloseWaite() {
    mini.unmask(document.body);
}
function Ajax(option) {
    OpenWaite();
    try {
        $.ajax({
            url: option.url,
            type: option.type,
            data: option.data,
            success: function (msg) {
                CloseWaite();
                option.success(msg);
            },
            error: function (q, m, e) {
                CloseWaite();
                if (option.error) {
                    option.error();
                }
                mini.alert(m);
            }
        });
    } catch (e) {
        CloseWaite();
        mini.alert(e.message);
    }
}
function CheckForm(form) {
    if (typeof form == "String") {
        form = new mini.Form(form);
    }
    form.validate();
    if (!form.isValid) {
        var err = form.getErrorTexts().join(";");
        if (err=="") {
            err = "数据验证不通过";
        }
        mini.alert(err);
        return false;
    }
    else
        return true;
}