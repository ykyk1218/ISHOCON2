<?php

require 'vendor/autoload.php';

Flight::route('/', function(){
    echo 'hello world!';
});

Flight::route('/aa', function(){
    echo 'hello world!aa';
});

Flight::start();
