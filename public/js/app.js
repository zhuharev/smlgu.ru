window.b=function(f,h){function c(){var d=g.length;if(0<d)for(var a=0;a<d;a++){var k=g[a],e;if(e=k)e=k.getBoundingClientRect(),e=(0<=e.top&&0<=e.left&&e.top)<=(window.innerHeight||h.documentElement.clientHeight)+200;e&&(k.src=k.getAttribute("data-src"),g.splice(a,1),d=g.length,a--)}else h.removeEventListener?f.removeEventListener("scroll",c):f.detachEvent("onscroll",c)}var g=[];return{a:function(){for(var d=h.querySelectorAll("[data-src]"),a=0;a<d.length;a++)d[a].src="",g.push(d[a]);
c();h.addEventListener?(f.addEventListener("scroll",c,!1),f.addEventListener("load",c,!1)):(f.attachEvent("onscroll",c),f.attachEvent("onload",c))},c:c}}(this,document);b.a();

!function(a,c,e,d,f){this.$=function(b){return new $[e].i(b)};f={length:0,i:function(b){c.push.apply(this,b&&b.nodeType?[b]:""+b===b?c.slice.call(a.querySelectorAll(b)):/^f/.test(typeof b)?$(a).r(b):null)},r:function(b){/c/.test(a.readyState)?b():$(a).on("DOMContentLoaded",b);return this},on:function(b,c){return this.each(function(a){a["add"+d](b,c)})},off:function(b,c){return this.each(function(a){a["remove"+d](b,c)})},each:function(b,a){c.forEach.call(a=this,b);return a},text:function(b){return b===
[]._?this[0].textContent:this[0].textContent=b},rmClass:function(b){return this[0].className=this[0].className.replace(RegExp("\\b"+b+"\\b"),"")},addClass:function(b){this[0].className=this[0].className+" "+b+"";return this},hasClass:function(cN){el=this[0];if(el.classList)return el.classList.contains(cN);else return new RegExp('(^| )' + cN + '( |$)', 'gi').test(el.className);}};
$[e]=f.i[e]=f}(document,[],"prototype","EventListener");

$.ajax=function(a,c,e){if(this.r=new XMLHttpRequest){var d=this.r;e=e||"";d.onreadystatechange=function(){4==d.readyState&&200==d.status&&c(d.responseText)};""!==e?(d.open("POST",a,!0),d.setRequestHeader("Content-type","application/x-www-form-urlencoded")):d.open("GET",a,!0);d.setRequestHeader("X-Requested-With","XMLHttpRequest");d.send(e)}};

$.prototype.hide=function(){
	return this.each(function(b){
		b.style.display = 'none'
	})
}

$.prototype.show=function(){
	return this.each(function(b){
		b.style.display = ''
	})
}

$.param = function(obj, prefix) {
  var str = [];
  for(var p in obj) {
    var k = prefix ? prefix + "[" + p + "]" : p, v = obj[p];
    str.push(typeof v == "object" ? $.param(v, k) : encodeURIComponent(k) + "=" + encodeURIComponent(v));
  }
  return str.join("&");
};

$("a.vote-link").on("click",votelistener)

function votelistener(e) {
	e.preventDefault();
	$("a.vote-link").off("click",votelistener)
	var url = this.href + "&ajax=true"
	var node = this
	$(this).addClass("disabled")
	$.ajax(url,function(a){
		a=JSON.parse(a);
		if ("undefined"!=typeof a.error) {
			show_message(a.error,true)
			$(this).rmClass("disabled")
		}
			else {
				show_message(a.response,false)
			}
		$("a.vote-link").on("click",votelistener)
		var el = $(node.parentNode)
		el.addClass("c33 shadow-border text-center")
		var sum = a.Behind + a.Against
		el.text(sum)
		var parent = el[0].parentNode
		var bars = parent.querySelectorAll(".progress-bar")
		console.log(bars)
		bars[0].style.width = a.widthB + "%"
		bars[1].style.width = a.widthA + "%"
		$(bars[0]).text(a.Behind)
		$(bars[1]).text(a.Against)
	})
}

function show_message(m,error) {
	console.log(m)
}

var search = $(".search")
var input = $(".search input[type=search]")[0]
var list = $(".ajax-search-results")[0]
search.on("keyup", function(){
	var q = input.value
	if (q === "" ){
		return
	}
	q = parseLatin(q)
	$.ajax("/prepodavateli/poisk?ajax=true&q="+q, function(r){
	  var listcontent = ""
      var json = JSON.parse(r)
      if (json.length >5)
      	length = 5
      else
      	length = json.length
      for (i = 0; i < length; i++) {
        listcontent += "<li><a href=\"/prepodavateli/"+json[i].Slug+"\"><img height='40px' src='"+json[i].Img+"'>"+json[i].PrettyName+"</a>"
      }
      list.innerHTML = listcontent
	})
})

function parseLatin(text){
  var outtext = text;
  var lat2 = 'F<DULT~:PBQRKVYJGHCNEA{WXIOMS}">Zf,dult`;pbqrkvyjghcnea[wxioms]\'.z';
  var rus2 = 'АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЬЫЪЭЮЯабвгдеёжзийклмнопрстуфхцчшщьыъэюя';
  for (var i = 0, l = lat2.length; i < l; i++) {
    outtext = outtext.split(lat2.charAt(i)).join(rus2.charAt(i));
  }
  return outtext;
}


$(".to-top").hide()

function totop() {

  window.scrollTo(0,100);

  var offset_top = 100

  function frame() {
    
    offset_top = offset_top - 10  // update parameters 
    
    window.scrollBy(0,-10); // show frame 

    if (offset_top == 0)  // check finish condition
      clearInterval(id)
  }

  var id = setInterval(frame, 10) // draw every 10ms
}

  $(".to-top").on("click",function(){
    totop()
  })

  window.onscroll=function(){
    var a=Math.round(window.pageYOffset||document.documentElement.scrollTop);
    650>=a&&($(".to-top").hide());
    650<a&&($(".to-top").show())
  };


$("#libsearch-val").on("change",function(){
	var val = $("#libsearch-val")[0].value
	$("#lib-search-form")[0].action=getRuslanLink(val)
})

