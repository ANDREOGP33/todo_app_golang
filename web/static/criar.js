$('#tform').on('submit', function(e) {
    e.preventDefault();

    $.ajax({
        url: '/criar',
        method: 'POST',
        data: {
            nome: $('#taskName').val(),
            descricao: $('#taskDescription').val()
        },
        success: function() {
            window.location.reload(); 
        },
        error: function(xhr) {
            console.error('Erro:', xhr.responseText);
        }
    });
});