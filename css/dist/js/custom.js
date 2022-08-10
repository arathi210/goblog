$(document).ready(function() {
    $('#example1').DataTable( {
        "paging":   false,        
        "info":     false,
        columnDefs: [
  { targets: 'no-sort', orderable: false }
]
    } );
} );