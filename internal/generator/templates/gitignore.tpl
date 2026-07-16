# ---- Go -------------------------------------------------------
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out
/vendor/

# ---- Environment ----------------------------------------------
.env
.env.*
!.env.example

# ---- Logs & storage -------------------------------------------
logs/
storage/

# ---- IDEs / OS ------------------------------------------------
.idea/
.vscode/
.DS_Store
Thumbs.db

# ---- Docker ---------------------------------------------------
docker-compose.override.yml

# ---- Build artifacts ------------------------------------------
dist/
tmp/