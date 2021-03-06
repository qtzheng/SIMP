var request = {
    QueryString: function(val) {
        var uri = window.location.search;
        var re = new RegExp("" + val + "=([^&?]*)", "ig");
        return ((uri.match(re)) ? (uri.match(re)[0].substr(val.length + 1)) : null);
    }
}
var IsOpenWindow = function() {
    var winId = request.QueryString("_winid");
    if (!winId && winId != "") {
        return true;
    } else
        return false;
}

function OpenWindow(url, title, iconCls, width, height, data, ondestroy) {
    mini.open({
        url: url, //页面地址
        title: title, //标题
        iconCls: iconCls, //标题图标
        width: width, //宽度
        height: height, //高度
        allowResize: true, //允许尺寸调节
        allowDrag: true, //允许拖拽位置
        showCloseButton: true, //显示关闭按钮
        showMaxButton: false, //显示最大化按钮
        showModal: true, //显示遮罩
        loadOnRefresh: true, //true每次刷新都激发onload事件
        onload: function() { //弹出页面加载完成
            if (typeof data != "undefined") {
                var iframe = this.getIFrameEl();
                //调用弹出页面方法进行初始化
                iframe.contentWindow.InitWindow(data);
            }
        },
        ondestroy: function(data) { //弹出页面关闭前
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
    win.set({
        title: title,
        iconCls: icons
    });
    win.showAtPos("center", "middle");
}

function OpenEditForm(url, form, win, title, icons, funcBefore, funcLoad) {
    OpenWaite();
    try {
        form.clear();
        if (typeof funcBefore == "function") {
            funcBefore();
        }
        $.ajax({
            url: url,
            type: "get",
            success: function(text) {
                CloseWaite();
                var data
                if (typeof text == "string") {
                    data = mini.decode(text); //反序列化成对象
                } else {
                    data = text;
                }
                if (typeof data.Result != "undefined" && data.Result == 0) {
                    form.setData(data.Message); //设置多个控件数据
                    if (typeof funcLoad == "function") {
                        funcLoad(data.Message);
                    }
                    win.set({
                        title: title,
                        iconCls: icons
                    });
                    win.showAtPos("center", "middle");
                } else {
                    mini.alert("系统出错！" + data.Message);
                }

            },
            error: function(q, m, e) {
                CloseWaite();
                mini.alert(m);
            }
        });
    } catch (e) {
        CloseWaite();
        mini.alert(e.message);
    }
}

function HideWin(win) {
    if (typeof win == "string") {
        mini.get(win).hide();
    } else
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
            success: function(msg) {
                CloseWaite();
                option.success(msg);
            },
            error: function(q, m, e) {
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
    try {
        if (typeof form == "string") {
            form = new mini.Form(form);
        }
        form.validate();
        if (!form.isValid()) {
            var err = form.getErrorTexts().join(";");
            if (err == "") {
                err = "数据验证不通过";
            }
            mini.alert(err);
            return false;
        } else
            return true;
    } catch (e) {
        mini.alert(e.Message);
        return false;
    }
}

function ShowTips(msg, time) {
    if (time) {
        time = 2000;
    };
    mini.showTips({
        content: msg,
        state: "default", //default|success|info|warning|danger
        x: "center", //left|center|right
        y: "center", //top|center|bottom
        timeout: time //自动消失间隔时间。默认2000（2秒）。});
    })
}

function SetAsLabel(c) {
    if (typeof c == "string") {
        c = mini.get(c);
    };
    c.setReadOnly(true); //只读
    c.setIsValid(true); //去除错误提示
    c.addCls("asLabel"); //增加asLabel外观
}

function SetAsInput(c) {
    if (typeof c == "string") {
        c = mini.get(c);
    };
    c.setReadOnly(false);
    c.removeCls("asLabel");
}