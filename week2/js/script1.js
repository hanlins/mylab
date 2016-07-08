
/** first relax the domain */
document.domain='mylab.com';
/** jquery: try to fetch an untrusted script, if success then alert to prompt. */
$.getScript("http://sub2.mylab.com/multi_csp_js/evil1.js", function(){
     alert("Script loaded but not necessarily executed.");
});
