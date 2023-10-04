import React, {useState, useEffect} from 'react';
import axios from "axios";
import {Button, Card, Row, Col, Form, Container, Modal } from 'react-bootstrap'
import 'bootstrap/dist/css/bootstrap.css'; 
const Task = ({taskData, setUpdateTaskName, deleteTask, setUpdateTask}) => {
    return (
        <Card>
            <Row>
                <Col>TaskName:{taskData!==undefined && taskData.task_name}</Col>
                <Col>TaskDetail:{taskData!==undefined && taskData.task_detail}</Col>
                <Col>TaskDate:{taskData!==undefined && taskData.date}</Col>
                <Col><Button onClick={() => deleteTask(taskData._id)}>delete task</Button></Col>
                <Col><Button onClick={() => updateTaskName()}>change waiter</Button></Col>
                <Col><Button onClick={() => updateTask()}>change order</Button></Col>
            </Row>
        </Card>
    )

    function updateTaskName(){
        setUpdateTaskName({
            "change": true,
            "id": taskData._id
        })
    }
    function updateTask(){
        setUpdateTask({
            "change": true,
            "id": taskData._id
        })
    }
}
export default Task