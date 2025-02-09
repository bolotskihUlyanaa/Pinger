import React, { useState, useEffect } from "react";
import axios from "axios";
import './client.css';

// Функция для отправки запроса на сервер
// и представления полученных данных в виде таблицы 
function Client() {
    // Создание переменных состояния и функций для их обновления
    const [data, setData] = useState([]); // Данные полученные от сервера
    const [error, setError] = useState(null); // Информация об ошибке
    useEffect(() => {
        // Определение асинхронной функции fetchData
        const fetchData = async () => {
            try {
                // GET запрос к серверу (ждет ответа сервера)
                const response = await axios.get(process.env.REACT_APP_API_URL);
                // Ответ сервера присваивается переменной data
                setData(response.data);
            } catch (err) {
                // Ошибка присваивается переменной error
                setError(err.message || "request error");
            }
        };
        fetchData(); //Вызов функции
    }, []);
    if (error) {
        return <h1>Error: {error}</h1>;
    }
    if (!data){
        return <h1>Loading...</h1>;
    }
    // Преобразование массива data в строки таблицы
    let res = data.map(function(item) {
        return <tr>
            <td>{item.Ip}</td>
            <td>{item.Time}</td>
            <td>{item.Date}</td>
            </tr>;
    });
    // Формирование HTML таблицы, в которую динамически добавляются строки
    return <table className="table">
        <thead>
            <tr className="column">
                <th>Ip</th>
                <th>Time, ns</th>
                <th>Date</th>
            </tr>
        </thead>
        <tbody>
            {res}
        </tbody>
    </table>;
}
export default Client;