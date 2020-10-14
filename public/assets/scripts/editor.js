

function loadCSS() {
    var $head = $("#preview").contents().find("head");
    $head.html("<style>" + editor2.getValue() + "</style>");
};

function loadJS() {
    var scriptTag = "<script>" + editorJS.getValue() + "<";
    scriptTag += "/script>";
    var previewFrame2 = document.getElementById('preview');
    var preview2 = previewFrame2.contentDocument || previewFrame2.contentWindow.document;
    preview2.open();
    preview2.write(editor.getValue() + scriptTag);
    preview2.close();
    loadCSS();
};

var delay;
// Initialize CodeMirror editor with a nice html5 canvas demo.
var editor = CodeMirror.fromTextArea(document.getElementById('code'), {
    mode: 'text/html',
    theme: "dracula",
    lineNumbers: true,
    styleActiveLine: true
});
editor.on("change", function () {
    clearTimeout(delay);
    delay = setTimeout(updatePreview, 300);
});

function updatePreview() {
    loadCSS();
}
setTimeout(updatePreview, 300);
var delay2;
// Initialize CodeMirror editor with a nice html5 canvas demo.
var editor2 = CodeMirror.fromTextArea(document.getElementById('codert'), {
    lineNumbers: true,
    styleActiveLine: true,
    matchBrackets: true,
    mode: 'css',
    theme: "dracula",
    gutters: ["CodeMirror-lint-markers"],
    lint: true
});
editor2.on("change", function () {
    clearTimeout(delay2);
    delay2 = setTimeout(updatePreview2, 300);
});

function updatePreview2() {
    var scriptTag = "<script>" + editorJS.getValue() + "<";
    scriptTag += "/script>";
    var previewFrame2 = document.getElementById('preview');
    var preview2 = previewFrame2.contentDocument || previewFrame2.contentWindow.document;
    preview2.open();
    preview2.write(editor2.getValue() + scriptTag);
    preview2.close();
    loadCSS();
}
setTimeout(updatePreview2, 300);


var delayJS;
// Initialize CodeMirror editor with a nice html5 canvas demo.
var editorJS = CodeMirror.fromTextArea(document.getElementById('code-js'), {
    lineNumbers: true,
    styleActiveLine: true,
    matchBrackets: true,
    mode: 'text/javascript',
    theme: "dracula",
    mode: "javascript",
    gutters: ["CodeMirror-lint-markers"],
    lint: true,
    lineWrapping: true,

});


editorJS.on("change", function () {
    clearTimeout(delayJS);
    delayJS = setTimeout(updatePreviewJS, 300);
});

function updatePreviewJS() {
    loadJS();
}
setTimeout(updatePreviewJS, 300);
//prints out html,css,js to be used as endpoints
function Save() {
    var html = editor.getValue();//http://codemirror.net/doc/manual.html#getValue
    var css = editor2.getValue();
    var js = editorJS.getValue();
    h=html.toString();
    cs=css.toString();
    j=js.toString();
    console.log("save running");
    console.group([h, cs, j]);
    console.groupEnd()
    
}


Save();