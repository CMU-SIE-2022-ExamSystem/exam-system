import React from 'react';
import {Container} from "react-bootstrap";

type Props = {
    children?: JSX.Element
}

function AppLayout({children} : Props) {
    return (
        <Container className="my-3 d-flex flex-column text-center">
            {children}
        </Container>
    );
}

export default AppLayout;