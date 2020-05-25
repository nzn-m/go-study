$(function () {
  $("#button").click(function () {
    var url = $("#url_post").val();

    var JSONdata = {
      email: $("#email").val(),
      password: $("#password").val(),
    };

    alert(JSON.stringify(JSONdata));

    $.ajax({
      type: "post",
      url: url,
      data: JSON.stringify(JSONdata),
      contentType: "application/JSON",
      dataType: "JSON",
      scriptCharset: "utf-8",
    })
      .success(function (data) {
        alert("success!!");
      })
      .error(function (XMLHttpRequest, textStatus, errorThrown) {
        alert("error!!!");
        console.log("XMLHttpRequest : " + XMLHttpRequest.status);
        console.log("textStatus     : " + textStatus);
        console.log("errorThrown    : " + errorThrown.message);
      });
  });
});
