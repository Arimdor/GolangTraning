(function () {
    let scene = new THREE.Scene();
    const aspectRatio = window.innerWidth / window.innerHeight;
    let camera = new THREE.PerspectiveCamera(75, aspectRatio, 0.1, 100);
    let render = new THREE.WebGLRenderer();

    render.setSize(window.innerWidth, window.innerHeight);
    document.body.appendChild(render.domElement);

    camera.position.x = 0.2;
    camera.position.y = 2;
    camera.position.z = 60;

    var moon;
    let loader = new THREE.TextureLoader();
    loader.load('../img/texture.jpg', function(texture){
        let geometry = new THREE.SphereGeometry(20,100,100);
        let material = new THREE.MeshBasicMaterial({
            map: texture
        });
        
        moon = new THREE.Mesh(geometry,material);
        moon.position.y = 10;
        scene.add(moon);
    });
    setTimeout(function(){ console.log("time gg")}, 1000);
    // let geometry = new THREE.BoxGeometry(10,10,10);
    let groundMaterial = new THREE.MeshPhongMaterial({
        color: 0xe6e6e6
    });
    // let mesh = new THREE.Mesh(geometry, groundMaterial);
    
    let pointLight = new THREE.PointLight(0xdfebff);
    pointLight.position.y = 10;
    pointLight.position.z = 90;

    scene.add(new THREE.AmbientLight(0x404040));
    scene.background = new THREE.Color(0xe6e6e6)
    // scene.add(mesh);
    scene.add(pointLight);

    function loop() {
        requestAnimationFrame(loop);
        moon.rotation.y += 0.020;
        moon.rotation.z += 0.008;
        render.render(scene, camera);
        console.log("fps");
    }
    loop();
})();