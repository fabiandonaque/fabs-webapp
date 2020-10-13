'use strict';

///////////////
//  Imports  //
///////////////

import { postJSON } from 'https://resources.fabsrobotics.com/js/utilities-v0.1.0.js';
import { DetailView } from 'https://resources.fabsrobotics.com/js/modules/fabsModules-master.js';

/////////////////////////
//  Window  listeners  //
/////////////////////////

window.addEventListener('load', () => {
	postJSON({width:window.innerWidth,height:window.innerHeight},"/",()=>{});
	setWebApp();
});

window.addEventListener('resize', e => {
	/*console.log("window resize");
	console.log(window.innerHeight);
	console.log(window.innerWidth);*/
});

function setWebApp(){
	let context = document.createElement("DIV");
	context.style = `
		width: 100%;
		height: 100%;
		background-color: aqua;
	`;
	let detailView = new DetailView();
	context.appendChild(detailView);
	document.body.appendChild(context);
}
