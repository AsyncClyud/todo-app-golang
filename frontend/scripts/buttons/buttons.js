"use strict";

function ShowPostForm() {
    var main_div = document.getElementById("main")
    main_div.innerHTML = ""

    var PostForm = document.createElement("div")

    PostForm.innerHTML = `
    <h2>Task Name</h2> <br>
    <input id="task_name"></input> <br>
    <h2>Description></h2> <br>
    <input id="description"></input> <br>
    <button onclick="Add_Tasks()">Post</button>
    `

    main_div.appendChild(PostForm)

}