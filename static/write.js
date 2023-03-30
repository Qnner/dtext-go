function uploadContent() {

    // If textarea value changes.
    if (content !== textarea.value) {
        var temp = textarea.value;
        var request = new XMLHttpRequest();

        request.open('POST', window.location.href, true);
        console.log(window.location.href)
        request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded; charset=UTF-8');
        request.onload = function() {
            if (request.readyState === 4) {

                // Request has ended, check again after 1 second.
                content = temp;
                setTimeout(uploadContent, 1000);
            }
        }
        request.onerror = function() {

            // Try again after 1 second.
            setTimeout(uploadContent, 1000);
        }
        request.send('msg=' + encodeURIComponent(temp));
    }
    else {

        // Content has not changed, check again after 1 second.
        setTimeout(uploadContent, 1000);
    }
}

var textarea = document.getElementsByClassName('content')[0];
var content = textarea.value;
textarea.onkeydown =  () => {
  if(event.code !== "Tab") return true;
  console.log('on key down: ', event.code)
  event.preventDefault();

  let start = this.selectionStart;
  let end = this.selectionEnd;
  if(start === end){
    document.execCommand('insertText',false,"  ");
  }
  else{
    let strBefore = this.value.slice(0,start);
    let curLineStart = strBefore.includes('\n')?strBefore.lastIndexOf('\n')+1 : 0;
    let strBetween = this.value.slice(curLineStart,end+1);
    let newStr = "  " + strBetween.replace(/\n/g,'\n  ');
    let lineBreakCount = strBetween.split('\n').length;
    let newStart = start + 2;
    let newEnd = end + (lineBreakCount + 1)*2;

    this.setSelectionRange(curLineStart,end);
    document.execCommand("insertText",false,newStr);
    this.setSelectionRange(newStart,newEnd);
  }
}

// Make the content available to print.
// printable.appendChild(document.createTextNode(content));

textarea.focus();
uploadContent();