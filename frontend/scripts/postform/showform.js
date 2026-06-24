"use strict";

function ShowPostForm() {
    const main_div = document.getElementById("main")
    main_div.innerHTML = ""

    var PostForm = document.createElement("div")
    PostForm.setAttribute("class", "sumbit_form")

    PostForm.innerHTML = `
    <h2>Task Name</h2> <br>
    <input id="task_name"></input> <br>
    <h2>Description></h2> <br>
    <input id="description"></input> <br>
    <button onclick="Add_Task()">Sumbit</button>
    `

    main_div.appendChild(PostForm)

}

function ShowEditForm(id) {
    const main_element = document.getElementById("main")
    
    const task_id = Number(document.getElementById(`Id_${id}`).textContent)
    var task_name = document.getElementById(`Task_Name_${id}`).textContent
    var description = document.getElementById(`Description_${id}`).textContent

    console.log(`${task_id}, ${task_name}, ${description}`)

    var EditForm = document.createElement("div")
    EditForm.setAttribute("class", "sumbit_form")

    EditForm.innerHTML = `
        <h2>Task Name</h2>
        <p>${task_name}</p>
        <input id="task_name"></input> <br>
        <h2>Description</h2>
        <p>${description}</p>
        <input id="description"></input> <br>
        <button onclick="Update_Task(${task_id})">Sumbit</button>
        `
    main_element.innerHTML = ""
    main_element.appendChild(EditForm)
}

async function GetIdAndShowEditForm() {
    const main_div = document.getElementById("main")
    var handleclick = function(event) {
        const task = document.getElementById(event.target.parentNode.id)
        const index = Number(String(task.id).indexOf("_") + 1) // + 1 для того, чтобы выводил число без _
        const id = String(task.id).substring(index)
        ShowEditForm(id)
        main_div.removeEventListener('click', handleclick(event))
    }
    main_div.addEventListener('click', handleclick(event))

}