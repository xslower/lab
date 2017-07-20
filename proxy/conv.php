<?php
if($argc < 2){
	echo '""';
	exit;
}
$path = $argv[1];
if(!file_exists($path)){
	echo "$path is not exist";
	exit;
}
$config = require($path);
$nodes = $config['nodes'];
foreach ($nodes as $i => $nod) {
	foreach ($nod as $j => $v) {
        // $nodes[$i][$j]['user'] = $v['username'];
		if(is_int($v)) $nodes[$i][$j] = strval($v);
	}
}
$parsedNodes = [];
foreach ($nodes as $k => $v) {
    if (!isset($v['role'])) {
        $v['role'] = 'master';
    }
    switch ($v['role']) {
        case 'master':
            $parsedNodes[$k][$k] = $v;
            break;
        case 'slave':
        case 'standby':
            if (isset($v['master']) && isset($nodes[$v['master']])) {
                //has key 'master'，and it's valid
                $parsedNodes[$v['master']][$k] = $v;
            } else {
                throw new \Exception($v['role'] . ' specified a not valid master value:' . $v['master']);
            }
            break;
        default:
            throw new \Exception('Not supported role:' . $v['role']);
    }
}

$json = json_encode($parsedNodes);
echo $json;
?>