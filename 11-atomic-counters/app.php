<?php

$time = time();
$i = 0;

do {
    $i++;
} while(time() - $time < 12);

echo $i;
