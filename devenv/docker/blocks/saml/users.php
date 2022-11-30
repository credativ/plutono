<?php
$config = array(
    'admin' => array(
        'core:AdminPassword',
    ),
    'example-userpass' => array(
        'exampleauth:UserPass',
        'saml-admin:plutono' => array(
            'groups' => array('admins'),
            'email' => 'saml-admin@example.com',
        ),
        'saml-editor:plutono' => array(
            'groups' => array('editors'),
            'email' => 'saml-editor@example.com',
        ),
        'saml-viewer:plutono' => array(
            'groups' => array(),
            'email' => 'saml-viewer@example.com',
        ),
    ),
);
