import React from "react";
import * as Yup from "yup";
import { Formik, Field, Form, ErrorMessage } from "formik";
import { FaTimes } from "react-icons/fa";

import { InputMask } from "@react-input/mask";
import { ddMMYYYtoMMDDYYYY } from "../../helpers/date";
import { backend } from "../../backend";

const UpdateUser = ({
  users,
  updateUserIndex,
  setupdateUserIndex,
  setUsers,
}) => {
  const user = users[updateUserIndex];

  // Yup validation schema
  const validationSchema = Yup.object().shape({
    name: Yup.string()
      .min(3, "Name must be at least 3 characters")
      .required("Name is required"),
    email: Yup.string()
      .email("Invalid email address")
      .required("Email is required"),
    date_of_birth: Yup.string().required("Birth date is required"),
  });

  const handleSubmit = (values, { resetForm }) => {
    const payload = { ...values };

    payload.date_of_birth = ddMMYYYtoMMDDYYYY(payload.date_of_birth);

    backend
      .put(`http://localhost:8080/users/${user.id}`, payload)
      .then(() => {
        alert("User updated successfully!");
        setUsers((prevUsers) =>
          prevUsers.map((u) => (u.id === user.id ? { ...u, ...values } : u))
        );
        resetForm();
        handleClose();
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const handleClose = () => {
    setupdateUserIndex(-1); // Close form
  };

  return (
    <>
      {updateUserIndex !== -1 && (
        <div className="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-50 z-50">
          <div className="w-full max-w-lg p-6 rounded-lg shadow-lg relative overflow-auto bg-base-100">
            <button
              className="absolute right-4 top-4 text-gray-500 hover:text-gray-700"
              onClick={handleClose}
            >
              <FaTimes />
            </button>
            <h2 className="text-3xl font-bold mb-4 text-center text-primary">
              Update User
            </h2>
            <Formik
              initialValues={{
                name: user.name,
                email: user.email,
                date_of_birth: user.date_of_birth,
              }}
              validationSchema={validationSchema}
              onSubmit={handleSubmit}
            >
              {({ handleChange }) => (
                <Form noValidate>
                  <div className="grid grid-cols-1 gap-4">
                    <div className="form-control w-full">
                      <label className="label">
                        <span className="label-text text-secondary">Name</span>
                      </label>
                      <Field
                        type="text"
                        name="name"
                        className="input input-bordered w-full"
                        onChange={handleChange}
                      />
                      <ErrorMessage
                        name="name"
                        component="span"
                        className="text-red-500 text-xs"
                      />
                    </div>

                    <div className="form-control w-full">
                      <label className="label">
                        <span className="label-text text-secondary">Email</span>
                      </label>
                      <Field
                        type="email"
                        name="email"
                        className="input input-bordered w-full"
                        onChange={handleChange}
                      />
                      <ErrorMessage
                        name="email"
                        component="span"
                        className="text-red-500 text-xs"
                      />
                    </div>

                    <div className="form-control w-full">
                      <label className="label">
                        <span className="label-text text-secondary">
                          Birth Date
                        </span>
                      </label>

                      <Field name="date_of_birth">
                        {({
                          field, // { name, value, onChange, onBlur }
                        }) => (
                          <InputMask
                            className="input input-bordered w-full"
                            mask="dd-mm-yyyy"
                            replacement={{ d: /\d/, m: /\d/, y: /\d/ }}
                            placeholder="DD-MM-YYYY"
                            {...field}
                          />
                        )}
                      </Field>
                      <ErrorMessage
                        name="date_of_birth"
                        component="span"
                        className="text-red-500 text-xs"
                      />
                    </div>

                    <button
                      type="submit"
                      className="btn btn-primary w-full mt-4"
                    >
                      Update User
                    </button>
                  </div>
                </Form>
              )}
            </Formik>
          </div>
        </div>
      )}
    </>
  );
};

export default UpdateUser;
