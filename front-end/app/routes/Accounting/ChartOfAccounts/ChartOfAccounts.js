import React from 'react';
import {
    Container,
    Row,
    Col
} from './../../../components';
import { setupPage } from './../../../components/Layout/setupPage';

import { HeaderMain } from "../../components/HeaderMain";


const ChartOfAccounts = () => (
    <Container>
        <Row className="mb-2">
            <Col lg={ 12 }>
                <HeaderMain 
                    title="Chart of accounts"
                    className="mb-4 mb-lg-3"
                />

            </Col>
        </Row>
    </Container>
);

export default setupPage({
    pageTitle: 'ChartOfAccounts'
})(ChartOfAccounts);