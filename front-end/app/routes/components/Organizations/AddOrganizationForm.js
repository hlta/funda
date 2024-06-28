import React from 'react';
import { useForm } from 'react-hook-form';
import PropTypes from 'prop-types';

const AddOrganizationForm = ({ onSubmit }) => {
    const { register, handleSubmit, formState: { errors } } = useForm();

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <div className="form-group">
                <label htmlFor="businessName">Business Name</label>
                <input
                    type="text"
                    className="form-control"
                    id="businessName"
                    {...register('businessName', { required: true })}
                />
                {errors.businessName && <span className="text-danger">Business Name is required</span>}
            </div>
            <div className="form-group">
                <label htmlFor="industry">Industry</label>
                <input
                    type="text"
                    className="form-control"
                    id="industry"
                    {...register('industry')}
                />
            </div>
            <div className="form-group">
                <label>Are you registered for GST?</label>
                <div className="form-check">
                    <input
                        type="checkbox"
                        className="form-check-input"
                        id="gstRegistered"
                        {...register('gstRegistered')}
                    />
                    <label className="form-check-label" htmlFor="gstRegistered">
                        Yes, calculate GST on my transactions
                    </label>
                </div>
            </div>
            <button type="submit" className="btn btn-primary">Add Organization</button>
        </form>
    );
};

AddOrganizationForm.propTypes = {
    onSubmit: PropTypes.func.isRequired,
};

export default AddOrganizationForm;
