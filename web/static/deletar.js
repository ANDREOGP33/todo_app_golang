$('.delete-btn').on('click', function(e) {
    e.preventDefault();

    $.ajax({
        url: '/delete',
        method: 'POST',
        data: {
            nome: $(this).data("id"),
        },
        success: function() {
            window.location.reload();
        },
        error: function(xhr) {
            console.error('Erro:', xhr.responseText);
        }
    });
});