$(document).ready(function () {


    var table = $('#myTable').DataTable({
        cache: false,
        "processing": true,
        "serverSide": true,

        "aLengthMenu": [[5, 15, 25, 50, 100, 200, 500], ['5', '15', '25', '50', '100', '200', '500']],
        "iDisplayLength": 10,

        "ajax": {
            "url": "/ajax",
            "type": "POST"
        },

        "columns": [
            { "data": "nama", "data-id": "test" },
            { "data": "alamat" },
            { "data": "email" },
            {
                "data": null,
                "render": function (data, type, row, meta) {

                    return "<a href='javascript:void(0)' data-toggle='tooltip'  data-id=" + row.uid + " data-original-title='Edit' class='edit btn btn-info btn-sm editUser'>Edit</a><a href='javascript:void(0)' data-toggle='tooltip'  data-id=" + row.uid + " data-original-title='Delete' class='delete btn btn-danger btn-sm deleteUser'>Delete</a>"

                }
            },


        ],

    });

    $.ajaxSetup({
        headers: {
            'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
        }
    });

    //click edit user
    $('body').on('click', '.editUser', function (e) {
        e.preventDefault()
        $('#modelHeading').html('Edit')

        var currentRow = $(this).closest("tr");

        var col1 = currentRow.find("td:eq(0)").text(); // get current row 2nd TD
        var col2 = currentRow.find("td:eq(1)").text(); // get 
        var col3 = currentRow.find("td:eq(2)").text(); // get 
        var id = $(this).data("id")

        $('#ajaxModel').modal('show');
        $('#id').val(id);
        $('#formnama').val(col1);
        $('#formalamat').val(col2);
        $('#formemail').val(col3);
    });
    //  $('#tombolRefreshTable').click(function(){
    //   table.load();
    //   });

    // click add user
    $('body').on('click', '.addUser', function (e) {
        e.preventDefault()
        $('#userForm').trigger("reset");
        $('#modelHeading').html('Add')

        $('#ajaxModel').modal('show');


    });

    $('#saveBtn').click(function (e) {
        e.preventDefault();

        if ($('#modelHeading').html() == "Add") {

            //confirm add user
            $(this).html('Sending..');
            $.ajax({
                data: $('#userForm').serialize(),
                url: "/user/add",
                type: "POST",
                dataType: 'json',
                success: function (data) {
                    $('#userForm').trigger("reset");
                    $('#ajaxModel').modal('hide');
                    $('#saveBtn').html('Simpan');
                    $('#successeditdata').before(`<div class="alert alert-success alert-dismissible fade show" role="alert"> Sukses add data user!!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                    $('#tombolRefreshTable').trigger('click');

                    table.ajax.reload();
                    $('.alert').delay(1000).fadeOut(function () {
                        $(this).remove();
                    });
                },
                error: function (data) {
                    console.log('Error:', data);
                    $('#saveBtn').html('Save Changes');
                }
            });

        } else if ($('#modelHeading').html() == "Edit") {
            //confirm edit user

            $(this).html('Sending..');
            var id_user = $("#id").val();
            console.log(id_user)
            console.log("heeko")

            $.ajax({
                data: $('#userForm').serialize(),
                url: "/user/save",
                type: "PUT",
                dataType: 'json',
                success: function (data) {
                    $('#userForm').trigger("reset");
                    $('#ajaxModel').modal('hide');
                    $('#saveBtn').html('Simpan');
                    $('#successeditdata').before(`<div class="alert alert-success alert-dismissible fade show" role="alert"> Sukses edit data user!!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                    $('#tombolRefreshTable').trigger('click');

                    table.ajax.reload();
                    $('.alert').delay(1000).fadeOut(function () {
                        $(this).remove();
                    });
                },
                error: function (data) {
                    console.log('Error:', data);
                    $('#saveBtn').html('Save Changes');
                }
            });
        }




    });

    //delete
    $('body').on('click', '.deleteUser', function () {

        var id = $(this).data("id");
        var r = confirm("Are You sure want to delete !");
        if (r == true) {
            $.ajax({
                type: "DELETE",
                url: "/user" + '/' + id,
                success: function (data) {
                    $('.table-responsive').before(`<div class="alert alert-danger alert-dismissible fade show" role="alert"> Sukses hapus user!!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                    table.ajax.reload();
                    $('.alert').delay(2000).fadeOut(function () {
                        $(this).remove();
                    });
                },
                error: function (data) {
                    console.log('Error:', data);
                }
            });
        } else {
            table.ajax.reload();
        }
    });
})