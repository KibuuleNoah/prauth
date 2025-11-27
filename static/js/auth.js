$(()=>{
  authType = getQueryParam("authType") || "signup"

  // alert(getQueryParam("authType"))

  $("[data-auth-toggler]").click(function(){
    if (authType === "signup"){
      $("#signupSection").slideUp()
      $("#signinSection").slideDown()
      authType = "signin"
    }else{
      authType = "signup"
      $("#signinSection").slideUp()
      $("#signupSection").slideDown()
    }
  })
  $("[data-auth-toggler]").click()

  function getQueryParam(param) {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get(param);
  }

  $('#signupForm').submit(function(event) {
    if (!validateSignupForm()) {
      event.preventDefault();
    }
  });

  // validate Email
  $('[name="email"]').blur(function() {
    if ($(this).val().length <= 0){
      $(this).parent().addClass("input-error")
      $(this).next().text("Required")
    }if (isValidEmail($(this).val()) != true){
      $(this).parent().addClass("input-error")
      $(this).next().text("Check Again!! Invalid Email")
    }else{
      $(this).parent().removeClass("input-error")
      $(this).next().text("Nice!!!")
    }
  });

  $('[name="password"]').blur(function() {
    const { errCount, msg} = checkPasswordStrength($(this).val())
    $(this).next().html(msg)

    if (errCount > 0){
      $(this).parent().addClass("input-error")
    }else{
      $(this).parent().removeClass("input-error")
    }

  })

  $('#termsAndCond').change(function() {
    if ($(this).is(':checked')){
      $("#signupBtn").prop("disabled",false)
    }else{
      $("#signupBtn").prop("disabled",true)
    }
    
  });
  function isValidEmail(email) {
    // Basic RFC 5322 compliant regex (covers most common cases)
    const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return regex.test(email);
  }

  function checkPasswordStrength(password) {
    const errors = [];
    const hasLowercase = /[a-z]/.test(password);
    const hasUppercase = /[A-Z]/.test(password);
    const hasNumber = /\d/.test(password);
    const hasSpecialChar = /[^a-zA-Z0-9]/.test(password);
    const hasMinLength = password.length >= 8;

    if (!hasLowercase) errors.push('at least one lowercase letter');
    if (!hasUppercase) errors.push('at least one uppercase letter');
    if (!hasNumber) errors.push('at least one number');
    if (!hasSpecialChar) errors.push('at least one special character');
    if (!hasMinLength) errors.push('at least 8 characters');

    return {
      errCount: errors.length,
      msg: errors.length > 0 ? `Password is missing: <br/>${errors.join('<br/> ')}` : 'Password is strong'
    }
  }

})

