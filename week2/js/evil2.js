/** This evil script is from domain sub2.mylab.com, it will try to access its parent page's dom */
/** first step, relax domain */
document.domain='mylab.com';
//alert('executing!');
/** get its parent page's dom object */
var s1=window.parent.document.getElementById("slot1");
s1.innerHTML='Flag captured!';
