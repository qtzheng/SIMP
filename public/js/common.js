var request = {
    QueryString: function (val) {
        var uri = window.location.search;
        var re = new RegExp("" + val + "=([^&?]*)", "ig");
        return ((uri.match(re)) ? (uri.match(re)[0].substr(val.length + 1)) : null);
    }
}
var IsOpenWindow= function() {
    var winId = request.QueryString("_winid");
    if (!winId && winId != "") {
        return true;
    }
    else
        return false;
}
function OpenWindow(url, title, iconCls, width, height, data, ondestroy) {
    mini.open({
        url: url,        //ҳ���ַ
        title: title,      //����
        iconCls: iconCls,    //����ͼ��
        width: width,      //���
        height: height,     //�߶�
        allowResize: true,       //����ߴ����
        allowDrag: true,         //������קλ��
        showCloseButton: true,   //��ʾ�رհ�ť
        showMaxButton: false,     //��ʾ��󻯰�ť
        showModal: true,         //��ʾ����
        loadOnRefresh: true,       //trueÿ��ˢ�¶�����onload�¼�
        onload: function () {       //����ҳ��������
            if (typeof data != "undefined") {
                var iframe = this.getIFrameEl();
                //���õ���ҳ�淽�����г�ʼ��
                iframe.contentWindow.InitWindow(data);
            }
        },
        ondestroy: function (data) {  //����ҳ��ر�ǰ
            if (typeof ondestroy == "function") {
                data = mini.clone(data);
                ondestroy(data);
            }
        }
    });
}
function CloseWindow(action, data, form) {
    if (form.isChanged) {
        if (action == "close" && form.isChanged()) {
            if (confirm("���ݱ��޸��ˣ��Ƿ��ȱ��棿")) {
                return false;
            }
        }
    }
   
    if (window.CloseOwnerWindow&&IsOpenWindow())
        return window.CloseOwnerWindow(data);
    //else
    //    window.close();
}
function OpenAddForm(form, win, title, icons, funcBefore) {
    form.clear();
    if (typeof funcBefore =="function") {
        funcBefore();
    }
    win.set({ title: title, iconCls: icons });
    win.showAtPos("center", "middle");
}
function OpenEditForm(url, form, win, title, icons, func) {
    OpenWaite();
    formModule.clear();
    $.ajax({
        url: url,
        type: "get",
        success: function (text) {
            CloseWaite();
            var data = mini.decode(text);   //�����л��ɶ���
            formModule.setData(data);             //���ö���ؼ�����
            if (func) {
                func(data);
            }
        }, error: function (q, m, e) {
            CloseWaite();
            if (q.responseText && q.responseText != "") {
                mini.alert(q.responseText);
            }
            else
                mini.alert(m);
        }
    });
    win.set({ title: title, iconCls: icons });
    win.showAtPos("center", "middle");
}