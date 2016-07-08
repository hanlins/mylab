/** current loadded by domain mylab.com.
 *  CSP policy should only white list script source from domain sub1.com
 */
document.domain='mylab.com';
var s1 = document.getElementById("slot1");
var if1 = document.createElement("iframe");
if1.src = 'http://sub2.mylab.com/multi_csp/iframe2.html';
if1.height = 400;
if1.width = 400;
if1.id='if1';
//s1.style.display='none';
s1.appendChild(if1);// append bridge
