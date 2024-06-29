import React from 'react';
import { useForm } from 'react-hook-form';
import PropTypes from 'prop-types';

const AddOrganizationForm = ({ onSubmit }) => {
    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm();

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <div className="form-group">
                <label htmlFor="name">Business Name</label>
                <input
                    type="text"
                    className="form-control"
                    id="name"
                    {...register('name', {
                        required: 'Business Name is required',
                        maxLength: {
                            value: 50,
                            message: 'Business Name cannot exceed 50 characters'
                        }
                    })}
                />
                {errors.name && (
                    <span className="text-danger">{errors.name.message}</span>
                )}
            </div>
            <div className="form-group">
                <label htmlFor="industry">Industry</label>
                <input
                    type="text"
                    className="form-control"
                    id="industry"
                    {...register('industry', {
                        maxLength: {
                            value: 50,
                            message: 'Industry cannot exceed 50 characters'
                        }
                    })}
                />
                {errors.industry && (
                    <span className="text-danger">{errors.industry.message}</span>
                )}
            </div>
            <div className="form-group">
                <label>Are you registered for GST?</label>
                <div className="form-check">
                    <input
                        type="checkbox"
                        className="form-check-input"
                        id="gst_registered"
                        {...register('gst_registered')}
                    />
                    <label className="form-check-label" htmlFor="gst_registered">
                        Yes, calculate GST on my transactions
                    </label>
                </div>
            </div>
            <button type="submit" className="btn btn-primary">
                Add Organization
            </button>
        </form>
    );
};

AddOrganizationForm.propTypes = {
    onSubmit: PropTypes.func.isRequired,
};

export default AddOrganizationForm;
