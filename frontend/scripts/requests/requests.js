"use strict";

async function Get_Tasks() {
    const response = await fetch(
        "/get_tasks", {
            method: "GET"
    })

    if(response.ok){

    const tasks = JSON.parse(await response.json())

    var main_element = document.getElementById("main")
    main_element.textContent = ""

    tasks.forEach(task => {
        const TaskDiv = document.createElement("div")
        TaskDiv.setAttribute("class", "task")

        TaskDiv.innerHTML = `
        <h2 id="Id">Id ${task.Id}</h2>
        <p id="Task_Name"><strong>${task.Task_Name}</strong></p>
        <p id="Description">${task.Description}</p>
        `

        main_element.appendChild(TaskDiv)
    });

    }
}

async function Add_Tasks() {
    const task_name = document.getElementById("task_name").value
    const description = document.getElementById("description").value

    const response = await fetch("/add_tasks", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({
            Task_Name: task_name,
            Description: description
        })
    })
    if(response.ok){
        console.log("OK")
        document.getElementById("response").textContent = "OK"
    }
    else{
        console.log(response)
    }

}

