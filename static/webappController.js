/*
	Created by:		Fabián Doñaque
	Copyright by:	Fabs Robotics SLU
	Created on:		2020-11-04
*/

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
	//setWidthAndHeight();
	dpr = window.devicePixelRatio;
	postJSON({event:"load",innerHeight:window.innerHeight,innerWidth:window.innerWidth,visualHeight: window.visualViewport.height,visualWidth: window.visualViewport.width},"/api/size", response => {
		console.log(response);
	});
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
	setTimeout(() => { setWidthAndHeight();},500);
	postJSON({event:"resize",innerHeight:window.innerHeight,innerWidth:window.innerWidth,visualHeight: window.visualViewport.height,visualWidth: window.visualViewport.width, computedHeight: window.getComputedStyle(document.documentElement).getPropertyValue("height"),divHeight: window.getComputedStyle(document.getElementById("context")).getPropertyValue("height")},"/api/size", response => {
		console.log(response);
	});
});

window.addEventListener("orientationchange", () => {
	//setWidthAndHeight();
	//resizeWebApp();
	postJSON({event:"orientationchange",innerHeight:window.innerHeight,innerWidth:window.innerWidth,visualHeight: window.visualViewport.height,visualWidth: window.visualViewport.width},"/api/size", response => {
		console.log(response);
	});
});

/////////////////
//  Functions  //
/////////////////

function setWidthAndHeight(){
	let heightText = parseInt(window.getComputedStyle(document.documentElement).getPropertyValue("height").slice(0,-2),10);
	let widthText = parseInt(window.getComputedStyle(document.documentElement).getPropertyValue("width").slice(0,-2),10);
	if(window.innerHeight != heightText){
		document.documentElement.style.height = window.innerHeight+"px";
		document.body.style.height = window.innerHeight+"px";
	}
	if(window.innerWidth != widthText) document.documentElement.style.width = window.innerWidth+"px";
}

function setWebApp(){
	/*/ Contexto
	let context = document.createElement("DIV");
	context.id = "context";
	context.style = `
		width: 100%;
		height: 100%;
		background-color: blue;
	`;

	// skeleton
	let header = document.createElement("DIV");
	header.style.height = "75px";
	header.style.backgroundColor = "yellow";
	let contentView = document.createElement("DIV");
	contentView.style.height = "calc( 100% - 75px )";
	contentView.style.backgroundColor = "red";

	// ContentView
	let list = document.createElement("DIV");
	list.style= ` display: inline-block; width: 300px; height: 100%; background-color: green;`;
	let appView = document.createElement("DIV");
	appView.style = `display: inline-block;width: calc( 100% - 300px ); height: 100%; background-color: purple;`;
	
	// Appends to document
	contentView.appendChild(list);
	contentView.appendChild(appView);
	context.appendChild(header);
	context.appendChild(contentView);
	document.body.appendChild(context);*/
}

function resizeWebApp(){
	//document.getElementById("webAppHeader").size = Math.floor(height/2)+"px";
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
