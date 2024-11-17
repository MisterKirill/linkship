<h1 align="center">Linkship - All your links in one place 🚀</h1>

<p align="center">
    <em>Open-source link & profile sharing service.</em>
</p>

## Self-hosting

You can launch your own instance of Linkship by following this guide:

### Running backend

1. Make sure you have **Docker** installed
2. Go to the ```./backend```
3. Edit ```.env``` file and set ```JWT_SECRET```
4. Run ```docker compose up```

### Running frontend

1. Go to the ```./frontend```
2. Edit ```.env``` file and set ```VITE_BACKEND_URL```
3. To build frontend, run ```npm i``` and ```npm run build```
4. Host static files generated in ```./frontend/dist``` using any provider or a webserver (e.g. *Vercel*, *GitHub Pages* or *Nginx*)
