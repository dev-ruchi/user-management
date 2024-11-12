import React, { useState, useEffect } from "react";
import axios from "axios";
import { Link } from "react-router-dom";

const UserList = () => {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    fetchUsers();
  }, []);

  const fetchUsers = () => {
    axios.get(`http://localhost:8080/users`).then((response) => {
      setUsers(response.data);
      console.log(response.data);
    });
  };

  return (
    <div className="p-6">
      <h2 className="text-3xl font-bold text-gray-700">Users</h2>
      <table>
        <thead>
          <tr className="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
            <th className="py-3 px-3 text-left sticky left-0 bg-gray-200 z-10">
              Name
            </th>
            <th className="py-3 px-3 text-left">Email</th>
            <th className="py-3 px-3 text-left">BirthDate</th>
          </tr>
        </thead>
        <tbody className="text-gray-600 text-sm font-light">
          {users.map((user, index) => (
            <tr
              key={`${user.id}-${index}`}
              className={`border-b border-gray-200 hover:bg-gray-100 ${
                index % 2 === 0 ? "bg-gray-50" : "bg-white"
              }`}
            >
              <td className="py-3 px-3 text-left whitespace-nowrap sticky left-0 z-10">
                <Link
                  to={`/users/${user.id}`}
                  className="text-blue-500 hover:underline"
                >
                  {user.name}
                </Link>
              </td>
              <td className="py-3 px-3 text-left">{user.email}</td>
              <td className="py-3 px-3 text-left">{user.date_of_birth}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default UserList;
