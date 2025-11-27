$(()=>{

  form = {}

  svEarlyAccessForm = localStorage.getItem("svEarlyAccessForm")
  // check of form already exists
  if (!(!svEarlyAccessForm)){
    form = JSON.parse(svEarlyAccessForm)

    if (form.hasOwnProperty("id")){
      $("#formSection").hide()
      $("#reviewSection").hide()
      $('#cardSchoolName').text(form.schoolName)
      $('#cardSchoolType').text(form.schoolType)
      $('#cardSchoolLevel').text(form.schoolLevel)

      $("#thanksSection").show()
    }else{
      loadEaryAccessFormData(form)
      document.getElementById("joinBtn").innerText = "Continue"
    }
  }

  // VALIDATE FORM
  $('#schoolName').blur(function() {
    if ($(this).val().length < 10){
      $(this).parent().addClass("input-error")
      $(this).next().text("Check Again!! School Name Too Short")
    }else{
      $(this).parent().removeClass("input-error")
      $(this).next().text("")
    }
  });

  // validate Phone Number
  $('#phone').blur(function() {
    if (isValidUgandanPhoneNumber($(this).val()) != true){
      $(this).parent().addClass("input-error")
      $(this).next().text("Check Again!! Invalid Ugandan Number")
    }else{
      $(this).parent().removeClass("input-error")
      $(this).next().text("It will only be used to update you")
    }
  });

  // validate Email
  $('#email').blur(function() {
    if (isValidEmail($(this).val()) != true){
      $(this).parent().addClass("input-error")
      $(this).next().text("Check Again!! Invalid Email")
    }else{
      $(this).parent().removeClass("input-error")
      $(this).next().text("Used as alternative to update you")
    }
  });


  $("#joinBtn").click(function() {

    // Check if no field still has an error
    if ($(".input-error").length != 0){
      // stop progress
      return
    }

    for (requiredField of ["schoolName","phone","email"]){
      field = $(`#${requiredField}`)
      if (field.val().length == 0 ) {
        field.parent().addClass("input-error")
        field.next().text("Please, This is required")
        return
      }else{
        field.parent().removeClass("input-error")
        field.next().text("")
        form[requiredField] = field.val()
      }
    }
    form["schoolType"] = $("#schoolType").val()
    form["schoolLevel"] = $("#schoolLevel").val()

    localStorage.setItem("svEarlyAccessForm",JSON.stringify(form))
    $("#formSection").slideUp()
    $("#reviewSection").slideDown()

    $("#reviewDisplay").html(createReview(form))
  })

  $("#editformBtn").click(function() {
    $("#reviewSection").slideUp()
    $("#formSection").slideDown()
  })

  // $("#submitDetailsBtn").prop("disabled",false)

  $("#submitDetailsBtn").click(function(){

    const formData = new FormData();

    formData.append("data", JSON.stringify(form));

    $("#reviewSection").html("<h1 class='font-bold text-[var(--gold)]'>Submitting...</h1>")
    $.ajax({
      url: "/api/v1/testers", 
      method: "POST",
      data: formData,
      processData: false, // prevent jQuery from transforming data
      contentType: false, // let the browser set it (multipart/form-data)
      success: function (response) {
        form["id"] = response.testerId
        localStorage.setItem("svEarlyAccessForm",JSON.stringify(form))
        $("#formSection").hide()
        $("#reviewSection").hide()
        $('#cardSchoolName').text(form.schoolName)
        $('#cardSchoolType').text(form.schoolType)
        $('#cardSchoolLevel').text(form.schoolLevel)

      $("#thanksSection").show()

        console.log("Upload success:", response);
      },
      error: function (xhr) {
        console.error("Upload failed:", xhr.responseText);
      }
    });

  })
})

function generateTempPassword(length = 4) {
  const chars = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789";
  let pass = "";
  for (let i = 0; i < length; i++) {
    pass += chars.charAt(Math.floor(Math.random() * chars.length));
  }
  return pass;
}

function isValidEmail(email) {
  // Basic RFC 5322 compliant regex (covers most common cases)
  const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return regex.test(email);
}

function isValidUgandanPhoneNumber(phone) {
  // Remove all spaces and hyphens
  const cleaned = phone.replace(/[\s-]/g, '');

  // Define regex for various formats
  const regex = /^(?:\+256|256|0)?7[0-9]{8}$/;

  return regex.test(cleaned);
}


function createReview(form){
  reviewHtml = ""

  for (key in form){
    label = $(`#${key}`).prev().text()
    reviewHtml += `
          <div class="flex flex-col sm:flex-row sm:justify-between">
            <span class="font-medium text-gray-600">${label}:</span>
            <span class="text-[var(--gold)] font-semibold">${form[key]}</span>
          </div>
          `
  }
  return reviewHtml
}

function loadEaryAccessFormData(jsonObj) {
  // Loop through each key in the JSON object
  for (const key in jsonObj) {
    if (jsonObj.hasOwnProperty(key)) {
      const value = jsonObj[key];
      const $field = $("#" + key); // Find field by ID

      if ($field.length) {
        // Handle different field types
        if ($field.is("select")) {
          $field.val(value).trigger("change");
        } else if ($field.is(":checkbox") || $field.is(":radio")) {
          $field.prop("checked", !!value);
        } else {
          $field.val(value);
        }
      }
    }
  }
}
// cookie.js
const Cookie = {
  set: (name, value, days = 7, path = "/") => {
    const date = new Date();
    date.setTime(date.getTime() + days * 24 * 60 * 60 * 1000);
    const expires = "; expires=" + date.toUTCString();
    document.cookie = `${name}=${encodeURIComponent(value)}${expires}; path=${path}`;
  },

  get: (name) => {
    const cookies = document.cookie.split("; ");
    for (const cookie of cookies) {
      const [key, val] = cookie.split("=");
      if (key === name) return decodeURIComponent(val);
    }
    return null;
  },

  delete: (name, path = "/") => {
    document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=${path}`;
  },

  all: () => {
    const result = {};
    document.cookie.split("; ").forEach(cookie => {
      const [key, val] = cookie.split("=");
      if (key) result[key] = decodeURIComponent(val);
    });
    return result;
  }
};
