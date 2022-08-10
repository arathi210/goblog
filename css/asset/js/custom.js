//allow only numbers validation(restrict)
    function isNumber(evt) {
        evt = (evt) ? evt : window.event;
        var charCode = (evt.which) ? evt.which : evt.keyCode;
        if (charCode > 31 && (charCode < 48 || charCode > 57)) {
            return false;
        }
        return true;
    }

 $(document).ready(function() {
 // Advance filter submit
  $(window).keydown(function(event){
    if(event.keyCode == 13) {
      event.preventDefault();
      $("form").submit();
    }
  });
 // Advance filter submit ends...

// Advance filter showing when having get methods
// $.urlParam = function(name){
// 	var results = new RegExp('[\?&]' + name + '=([^&#]*)').exec(window.location.href);
// 	return results[1] || 0;
// }
// var type=$.urlParam('type'); 
// var sortby=$.urlParam('sortby');
// var beds=$.urlParam('beds');  
// var bath=$.urlParam('bath');  
// var min=$.urlParam('min'); 
// var max=$.urlParam('max');

// console.log(type);         
// 	if(type!="" || sortby!="" || beds!=""|| min!=""|| max!="" || bath!="")
// 	{
// 		$('#filter').show();
		
// 	}
// Advance filter showing when having get methods ends....
// Images Delete functionalities
 $(this).on("click",".remove_this",function() { 
  // alert();
  
        var rmv_div=$(this).parent().parent().attr('class');
        // alert(rmv_div);
        var str=rmv_div.split(" ");
        // alert(str[1]);
        var id=str[1].split("_");  
        // alert(id);      

        $("#x_"+id[1]).val("");
        $("#y_"+id[1]).val("");
        $("#w_"+id[1]).val("");
        $("#h_"+id[1]).val("");
        $("#imgstore_"+id[1]).val("");
  
        $("."+str[1]).remove();
   
  
  });
 // Summernote
    $('.textarea').summernote();

});

// Post fix
$(window).scroll(function(){
 if ( $(this).scrollTop() > 1500) {
       if($(window).height() > 800)
       {                       
        $('.pos_fix').addClass('fixed');             
        $("#big_resionR").css("display", "block");
        //$(".pos_fix").css("bottom", "1070px");    
       }
       else{
        $('.pos_fix').addClass('fixed');     
      }
 } 
 else
 {
     $('.pos_fix').removeClass('fixed');
 }

});
$(window).scroll(function() {
  if($(window).scrollTop() + $(window).height() == $(document).height())
  {
    //you are at bottom
    $('.pos_fix').addClass('endof_bottom');
  }
});
 $(window).on("scroll", function() {
  var scrollHeight = $(document).height();
  var scrollPosition = $(window).height() + $(window).scrollTop();
  
  if ((scrollHeight - scrollPosition) / scrollHeight === 0) {
       $('.pos_fix ').addClass('footer_positionunset')
       
  }
    else
    {
    $('.pos_fix ').removeClass('footer_positionunset')
    }
});
 


