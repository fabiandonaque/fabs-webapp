'use strict';

///////////////
//  Imports  //
///////////////

import { res,postJSON } from 'https://resources.fabsrobotics.com/js/utilities-v0.1.0.js';
import { SplitView } from 'https://resources.fabsrobotics.com/js/modules/fabsModules-master.js';

////////////////////////
//  Global Variables  //
////////////////////////

let touchable;
let width;
let height;
let dpr;
let treeApp;

/////////////////
//  Listeners  //
/////////////////

window.addEventListener('load', () => {
	touchable = "ontouchstart" in document.documentElement;
	setWidthAndHeight();
	dpr = window.devicePixelRatio;
	postJSON({},"/api/apps/getAppsTree", response => {
		if(response.code == res.Ok){
			treeApp = response.message;
			console.log(treeApp);
		}
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
	// Contexto
	let context = document.createElement("DIV");
	context.style = `
		width: var(--width);
		height: var(--height);
		background-color: blue;
	`;
	// skeleton
	let splitView = new SplitView();
	splitView.id = "webAppHeader";
	splitView.size = "100px";
	splitView.separator = true;
	let header = document.createElement("DIV");
	let contentView = document.createElement("DIV");
	header.style.backgroundColor = "yellow";
	// ContentView
	
	// Appends to document
	splitView.appendChild(header);
	splitView.appendChild(contentView);
	context.appendChild(splitView);
	document.body.appendChild(context);
}

function resizeWebApp(){
	document.getElementById("webAppHeader").size = Math.floor(height/2)+"px";
}

function getAppsForList(obj){
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
