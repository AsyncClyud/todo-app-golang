"use strict";


async function Get_Tasks() {
    const response = await fetch(
        "/get_tasks", {
            method: "GET"
    })

    if(response.ok){

    var numbering = 1
    const tasks = JSON.parse(await response.json())

    const main_element = document.getElementById("main")
    main_element.innerHTML = ""

    tasks.forEach(task => {
        const task_element = document.createElement("div")
        task_element.setAttribute("class", "task")
        task_element.setAttribute("id", `task_${numbering}`)

        task_element.innerHTML = `
        <h2 id="Id_${numbering}">${task.Id}</h2> 
        <p id="Task_Name_${numbering}"><strong>${task.Task_Name}</strong></p>
        <p id="Description_${numbering}">${task.Description}</p>
        <button onclick="GetIdAndShowEditForm()">update</button>
        <button onclick="GetIdAndDeleteTask()">delete</button>
        `

        numbering++;
        main_element.appendChild(task_element)
    });

    }
}

async function Add_Task() {
    const task_name = document.getElementById("task_name").value
    const description = document.getElementById("description").value

    const response = await fetch("/add_task", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({
            Task_Name: task_name,
            Description: description
        })
    })
    if(response.ok){
        document.getElementById("response").textContent = "OK"
    }
    else{
        console.log(response.status)
    }

}

async function Delete_Task(id) {

    response = await fetch("/delete_task", {
        method: "DELETE",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({
            Id: id
        })
    })

    if(response.ok){
        document.getElementById("response").textContent = "OK"
    }
    else{
        console.log(response.status)
    }

}

async function Update_Task(id) {
    const task_name = document.getElementById("task_name").value
    const description = document.getElementById("description").value

    response = await fetch("/update_task", {
        method: "PUT",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({
            Id: id,
            Task_Name: task_name,
            Description: description
        })
    })
    if(response.ok){
        document.getElementById("response").textContent = "OK"
    }
    else{
        console.log(response.status)
    }
}

async function GetIdAndDeleteTask() {
    const main_element = document.getElementById("main")
    var handleclick = function(event) {
        const task = document.getElementById(event.target.parentNode.id)
        const id = Number(document.getElementById(task.firstChild.nextSibling.id).textContent)
        main_element.removeChild(task)
        Delete_Task(id)
        main_element.removeEventListener('click', handleclick(event))
    }
    main_element.addEventListener('click', handleclick(event))
}

