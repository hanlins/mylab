/** XSS vulnability (type - 1)
 *
 *  In this function, it takes a user input and write that input directly to the webpage.
 *
 *  @return Void.
 */
function xss() {
  var payload = "<a href='http://www.baidu.com'><div style='color:red; background-color:green; opacity:0.3; position:absolute; top:30px; left:30px;'>test</div></a>";
  var input = prompt("Please enter the payload:", payload);
  document.write(payload);
}

/** Benigh behavior of clicking button.
 *
 *  @return Void.
 */
function good() {
  alert("I'm good guy!");
}

xss();
