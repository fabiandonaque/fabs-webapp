'use strict';

///////////////
//  Imports  //
///////////////

import { postJSON } from 'https://resources.fabsrobotics.com/js/utilities-v0.1.0.js';
import { SplitView } from 'https://resources.fabsrobotics.com/js/modules/fabsModules-master.js';

////////////////////////
//  Global Variables  //
////////////////////////

let touchable;
let width;
let height;
let dpr;

/////////////////
//  Listeners  //
/////////////////

window.addEventListener('load', () => {
	touchable = "ontouchstart" in document.documentElement;
	setWidthAndHeight();
	dpr = window.devicePixelRatio;
	postJSON({},"/api/apps/getAppsTree", response => {
		console.log(response);
	});
	setWebApp();
});

window.addEventListener('resize', e => {
	setWidthAndHeight();
	resizeWebApp();
	//postJSON({event:"resize",width:width,height:height,dpr:dpr,touchableDevice:touchable},"/",()=>{});
});

window.addEventListener("orientationchange", () => {
	setWidthAndHeight();
	resizeWebApp();
	/*postJSON({event:"orientationchange",height:height,width:width},"/", response => {
		console.log(response);
	});*/
});

/////////////////
//  Functions  //
/////////////////

function setWidthAndHeight(){
	width = window.innerWidth;
	height = window.innerHeight;
	document.documentElement.style.setProperty('--height', `${height}px`);
	document.documentElement.style.setProperty('--width', `${width}px`);
}

function setWebApp(){
	let context = document.createElement("DIV");
	context.style = `
		width: var(--width);
		height: var(--height);
		background-color: blue;
	`;
	let splitView = new SplitView();
	splitView.id = "webAppHeader";
	splitView.size = Math.floor(height/2)+"px";
	splitView.separator = true;
	let first = document.createElement("DIV");
	let second = document.createElement("DIV");
	first.style.backgroundColor = "yellow";
	splitView.appendChild(first);
	splitView.appendChild(second);
	context.appendChild(splitView);
	document.body.appendChild(context);


	// Control lines	
	let top = document.createElement("DIV");
	let bottom = document.createElement("DIV");
	top.style = `position:absolute;top:0;left:0;width: var(--width); height: 1px; background-color: red;`;
	bottom.style = `position:absolute;top:calc( var(--height) - 1px );left:0;width: var(--width); height: 1px; background-color: red;`;
	document.body.appendChild(top);
	document.body.appendChild(bottom);
}

function resizeWebApp(){
	document.getElementById("webAppHeader").size = Math.floor(height/2)+"px";
}

// iPhone Safari Ok
// iPhone Chrome Ok
// iPhone Firefox Ok
// iPad Safari Ok
// iPad Chrome Fail
// iPad Firefox Ok
// Android phone Samsung browser Ok
// Android phone Chrome Ok
// Android phone Firefox Ok
// Android tablet Samsung browser Ok
// Android tablet Chrome Ok
// Android tablet Firefox Ok
