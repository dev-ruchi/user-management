import React from "react";
import * as Yup from "yup";
import { Formik, Field, Form, ErrorMessage } from "formik";
import { FaTimes } from "react-icons/fa";
import { backend } from "../../backend";
import { InputMask } from "@react-input/mask";
import { ddMMYYYtoMMDDYYYY } from "../../helpers/date";
import dayjs from "dayjs";

const CreateUser = ({ createUser, setCreateUser, setUsers }) => {
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
      .post("http://localhost:8080/users", payload)
      .then(({ data: user }) => {
        user.date_of_birth = dayjs.utc(user.date_of_birth).format("DD-MM-YYYY");
        setUsers((prevUsers) => [...prevUsers, user]);
        resetForm();
        handleClose();
      })
      .catch((error) => {
        alert("Something went wrong");
        console.log(error);
      });
  };

  const handleClose = () => {
    setCreateUser(false); // Close form
  };

  return (
    <>
      {createUser && (
        <div className="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-50 z-50">
          <div className="w-full max-w-lg p-6 rounded-lg shadow-lg relative overflow-auto bg-base-100">
            <button
              className="absolute right-4 top-4 text-gray-500 hover:text-gray-700"
              onClick={handleClose}
            >
              <FaTimes />
            </button>
            <h2 className="text-3xl font-bold mb-4 text-center text-primary">
              Create New User
            </h2>
            <Formik
              initialValues={{
                name: "",
                email: "",
                date_of_birth: "",
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
                      Create User
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

export default CreateUser;
