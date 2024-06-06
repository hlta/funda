import React, { useState } from "react";
import { useForm, Controller } from "react-hook-form";
import { Link, useHistory } from "react-router-dom";
import { useAuth } from "../../../hooks/useAuth";

import {
  Form,
  FormGroup,
  FormFeedback,
  Input,
  CustomInput,
  Button,
  Label,
  EmptyLayout,
  ThemeConsumer,
  Alert,
} from "../../../components";

import { HeaderAuth } from "../../components/Pages/HeaderAuth";
import { FooterAuth } from "../../components/Pages/FooterAuth";

const Register = () => {
  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm();
  const { performRegister } = useAuth();
  const history = useHistory();
  const [apiError, setApiError] = useState("");

  const handleRegistration = async (data) => {
    try {
      console.log(data);
        await performRegister(data);
        history.push("/dashboard"); 
    } catch (error) {
        console.error("Registration failed:", error);
        setApiError("Failed to register. Please try again."); 
    }
};


  return (
    <EmptyLayout>
      <EmptyLayout.Section center width={480}>
        <HeaderAuth title="Create Account" />
        <Form className="mb-3" onSubmit={handleSubmit(handleRegistration)}>
          {apiError && (
            <Alert color="danger" className="mb-3">
              {apiError}
            </Alert> // Displaying the API error message
          )}
          <FormGroup>
            <Label for="firstName">First Name</Label>
            <Controller
              name="firstName"
              control={control}
              rules={{ required: "First name is required" }}
              render={({ field }) => (
                <Input
                  {...field}
                  type="text"
                  id="firstName"
                  placeholder="Enter your first name..."
                  className="bg-white"
                  invalid={errors.firstName ? true : false}
                />
              )}
            />
            {errors.firstName && (
              <FormFeedback>{errors.firstName.message}</FormFeedback>
            )}
          </FormGroup>
          <FormGroup>
            <Label for="lastName">Last Name</Label>
            <Controller
              name="lastName"
              control={control}
              rules={{ required: "Last name is required" }}
              render={({ field }) => (
                <Input
                  {...field}
                  type="text"
                  id="lastName"
                  placeholder="Enter your last name..."
                  className="bg-white"
                  invalid={errors.lastName ? true : false}
                />
              )}
            />
            {errors.lastName && (
              <FormFeedback>{errors.lastName.message}</FormFeedback>
            )}
          </FormGroup>
          <FormGroup>
            <Label for="emailAddress">Email Address</Label>
            <Controller
              name="email"
              control={control}
              rules={{ required: "Email is required" }}
              render={({ field }) => (
                <Input
                  {...field}
                  type="email"
                  id="emailAddress"
                  placeholder="Enter your email..."
                  className="bg-white"
                  invalid={errors.email ? true : false}
                />
              )}
            />
            {errors.email && (
              <FormFeedback>{errors.email.message}</FormFeedback>
            )}
          </FormGroup>
          <FormGroup>
            <Label for="password">Password</Label>
            <Controller
              name="password"
              control={control}
              rules={{ required: "Password is required" }}
              render={({ field }) => (
                <Input
                  {...field}
                  type="password"
                  id="password"
                  placeholder="Enter a password..."
                  className="bg-white"
                  invalid={errors.password ? true : false}
                />
              )}
            />
            {errors.password && (
              <FormFeedback>{errors.password.message}</FormFeedback>
            )}
          </FormGroup>
          <FormGroup>
            <Controller
              name="acceptTerms"
              control={control}
              rules={{
                required: "You must accept the terms and privacy policy",
              }}
              render={({ field }) => (
                <CustomInput
                  {...field}
                  type="checkbox"
                  id="acceptTerms"
                  label="Accept Terms and Privacy Policy"
                  inline
                  invalid={errors.acceptTerms ? true : false}
                />
              )}
            />
            {errors.acceptTerms && (
              <FormFeedback>{errors.acceptTerms.message}</FormFeedback>
            )}
          </FormGroup>
          <ThemeConsumer>
            {({ color }) => (
              <Button type="submit" color={color} block>
                Create Account
              </Button>
            )}
          </ThemeConsumer>
        </Form>
        <div className="d-flex mb-5">
          <Link to="/forgot-password" className="text-decoration-none">
            Forgot Password
          </Link>
          <Link to="/login" className="ml-auto text-decoration-none">
            Login
          </Link>
        </div>
        <FooterAuth />
      </EmptyLayout.Section>
    </EmptyLayout>
  );
};

export default Register;
