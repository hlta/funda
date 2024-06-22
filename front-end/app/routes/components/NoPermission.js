import React from 'react';
import { Container, Row, Col, Card, CardBody, CardTitle, CardText, Button } from 'reactstrap';

const NoPermission = () => (
    <Container className="mt-5">
        <Row className="justify-content-center">
            <Col md="8" lg="6">
                <Card className="text-center">
                    <CardBody>
                        <CardTitle tag="h4">Access Denied</CardTitle>
                        <CardText>You do not have permission to view this page.</CardText>
                        <Button color="primary" href="/">Go to Homepage</Button>
                    </CardBody>
                </Card>
            </Col>
        </Row>
    </Container>
);

export default NoPermission;
