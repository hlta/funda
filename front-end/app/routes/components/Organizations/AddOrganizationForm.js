import React, { useEffect } from 'react';
import { useForm, Controller } from 'react-hook-form';
import PropTypes from 'prop-types';
import {
    Input,
    CustomInput,
    FormFeedback
} from "../../../components";

const AddOrganizationForm = ({ onSubmit, serverErrors, clearServerErrors }) => {
    const {
        control,
        handleSubmit,
        setError,
        clearErrors,
        formState: { errors }
    } = useForm();

    useEffect(() => {
        // Set server errors in the form
        if (serverErrors) {
            for (const [field, message] of Object.entries(serverErrors)) {
                setError(field, { type: 'server', message });
            }
        }
    }, [serverErrors, setError]);

    const handleChange = (field) => {
        clearErrors(field); // Clear errors in react-hook-form
        clearServerErrors(field); // Clear errors in parent component state
    };

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <div className="form-group">
                <label htmlFor="organizationName">Business Name</label>
                <Controller
                    name="organizationName"
                    control={control}
                    rules={{
                        required: "Business name is required",
                        maxLength: {
                            value: 50,
                            message: "Business Name cannot exceed 50 characters"
                        }
                    }}
                    render={({ field }) => (
                        <Input
                            {...field}
                            type="text"
                            id="organizationName"
                            placeholder="Enter your business name..."
                            className="bg-white"
                            invalid={errors.organizationName ? true : false}
                            onChange={(e) => {
                                field.onChange(e);
                                handleChange('organizationName');
                            }}
                        />
                    )}
                />
                {errors.organizationName && (
                    <FormFeedback>{errors.organizationName.message}</FormFeedback>
                )}
            </div>
            <div className="form-group">
                <label htmlFor="industry">Industry</label>
                <Controller
                    name="industry"
                    control={control}
                    rules={{
                        maxLength: {
                            value: 50,
                            message: "Industry cannot exceed 50 characters"
                        }
                    }}
                    render={({ field }) => (
                        <Input
                            {...field}
                            type="text"
                            id="industry"
                            placeholder="Enter your industry..."
                            className="bg-white"
                            invalid={errors.industry ? true : false}
                            onChange={(e) => {
                                field.onChange(e);
                                handleChange('industry');
                            }}
                        />
                    )}
                />
                {errors.industry && (
                    <FormFeedback>{errors.industry.message}</FormFeedback>
                )}
            </div>
            <div className="form-group">
                <Controller
                    name="gst_registered"
                    control={control}
                    render={({ field }) => (
                        <CustomInput
                            {...field}
                            type="checkbox"
                            id="gst_registered"
                            label="Yes, calculate GST on my transactions"
                            inline
                            invalid={errors.gst_registered ? true : false}
                            onChange={(e) => {
                                field.onChange(e);
                                handleChange('gst_registered');
                            }}
                        />
                    )}
                />
            </div>
            <button type="submit" className="btn btn-primary">
                Add Organization
            </button>
        </form>
    );
};

AddOrganizationForm.propTypes = {
    onSubmit: PropTypes.func.isRequired,
    serverErrors: PropTypes.object,
    clearServerErrors: PropTypes.func.isRequired,
};

export default AddOrganizationForm;
