$(document).ready(function () {

    $.ajaxSetup({
        headers: {
            'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
        }
    });

    $('#btnReg').click(function (e) {
        e.preventDefault()

        console.log($('#registerForm').serialize())
        console.log("hello")
        //confirm add user
        $.ajax({
            data: $('#registerForm').serialize(),
            url: "/user/register",
            type: "POST",
            dataType: 'json',
            success: function (data) {
                // $('#registerForm').trigger("reset");

                window.location = "http://localhost:8080/user/login";

                // $('#loginlink').trigger('click');
                // $('#successregister').before(`<div class="alert alert-success alert-dismissible fade show" role="alert"> Sukses Register!!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                // console.log("data res")
                // $('.alert').delay(1000).fadeOut(function () {
                //     $(this).remove();
                // });
            },
            error: function (data) {
                console.log('Error:', data);
                $('#saveBtn').html('Save Changes');
            }
        });

    });

    //Login
    $('#btnLogin').click(function (e) {
        e.preventDefault()

        console.log($('#loginForm').serialize())
        console.log("hello")
        //confirm add user
        $.ajax({
            data: $('#loginForm').serialize(),
            url: "/user/login",
            type: "POST",
            dataType: 'json',
            success: function (data) {
                // $('#loginForm').trigger("reset");
                // alert("You will now be redirected.");
                window.location = "http://localhost:8080/";
                // $('#loginlink').trigger('click');
                // $('#successregister').before(`<div class="alert alert-success alert-dismissible fade show" role="alert"> Sukses Register!!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                // console.log("data res")
                // $('.alert').delay(1000).fadeOut(function () {
                //     $(this).remove();
                // });
            },
            error: function (data) {
                console.log('Error:', data);
                $('#saveBtn').html('Save Changes');
            }
        });

    });



})