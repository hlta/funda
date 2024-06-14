import React, { useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import { useForm, Controller } from 'react-hook-form';

import {
    Form,
    FormGroup,
    Input,
    Button,
    Label,
    EmptyLayout,
    ThemeConsumer,
    Alert,
    FormText,
    FormFeedback
} from './../../../components';

import { HeaderAuth } from "../../components/Pages/HeaderAuth";
import { FooterAuth } from "../../components/Pages/FooterAuth";

import { useAuth } from "../../../hooks/useAuth";

const Login = () => {
    const { control, handleSubmit, formState: { errors, isSubmitting } } = useForm();
    const { performLogin } = useAuth();
    const [globalError, setGlobalError] = useState('');
    const history = useHistory();

    const onSubmit = async (data) => {
        setGlobalError('');
        try {
            await performLogin(data);
            history.push('/');
        } catch (error) {
            if (error.response && error.response.status === 401) {
                setGlobalError('Invalid username or password.');
            } else {
                setGlobalError('An unexpected error occurred. Please try again later.');
            }
        }
    };

    return (
        <EmptyLayout>
            <EmptyLayout.Section center>
                { /* START Header */}
                <HeaderAuth 
                    title="Sign In to Funda"
                />
                { /* END Header */}
                { /* START Form */}
                <Form className="mb-3" onSubmit={handleSubmit(onSubmit)}>
                    {globalError && (
                        <Alert color="danger" className="mb-3">
                            {globalError}
                        </Alert>
                    )}
                    <FormGroup>
                        <Label for="email">
                            Email Address
                        </Label>
                        <Controller
                            name="email"
                            control={control}
                            rules={{ required: 'Email is required' }}
                            render={({ field }) => (
                                <Input
                                    {...field}
                                    type="email"
                                    id="email"
                                    placeholder="Enter email..."
                                    className="bg-white"
                                    invalid={!!errors.email}
                                />
                            )}
                        />
                        {errors.email && <FormFeedback>{errors.email.message}</FormFeedback>}
                    </FormGroup>
                    <FormGroup>
                        <Label for="password">
                            Password
                        </Label>
                        <Controller
                            name="password"
                            control={control}
                            rules={{ required: 'Password is required' }}
                            render={({ field }) => (
                                <Input
                                    {...field}
                                    type="password"
                                    id="password"
                                    placeholder="Password..."
                                    className="bg-white"
                                    invalid={!!errors.password}
                                />
                            )}
                        />
                        {errors.password && <FormFeedback>{errors.password.message}</FormFeedback>}
                    </FormGroup>
                    <ThemeConsumer>
                        {
                            ({ color }) => (
                                <Button color={ color } block type="submit" disabled={isSubmitting}>
                                    Sign In
                                </Button>
                            )
                        }
                    </ThemeConsumer>
                </Form>
                { /* END Form */}
                { /* START Bottom Links */}
                <div className="d-flex mb-5">
                    <Link to="/forgot-password" className="text-decoration-none">
                        Forgot Password
                    </Link>
                    <Link to="/register" className="ml-auto text-decoration-none">
                        Register
                    </Link>
                </div>
                { /* END Bottom Links */}
                { /* START Footer */}
                <FooterAuth />
                { /* END Footer */}
            </EmptyLayout.Section>
        </EmptyLayout>
    );
};

export default Login;
