{{define "form"}}
<!DOCTYPE html>
<html>

<head>
    {{template "_header" .}}
</head>

<body>
    
    {{template "_menu" .}}
    
    
    <!-- Begin page content -->
    <main class="flex-shrink-0 p-4">
        <form class="form form-horizontal" action="/task" method="post">
            <div class="row mb-3">
                <label for="Name" class="col-sm-2 col-form-label">Nama</label>
                <div class="col-sm-10">
                    <input type="text" name="name" class="form-control" required>
                </div>
            </div>
            <div class="row mb-3">
                <label for="Pic" class="col-sm-2 col-form-label">PIC</label>
                <div class="col-sm-10">
                    <input type="text" name="pic" class="form-control" required>
                </div>
            </div>
            <div class="row mb-3">
                <label for="Deadline" class="col-sm-2 col-form-label">Deadline</label>
                <div class="col-sm-10">
                    <input type="text" name="deadline" class="form-control" required>
                </div>
            </div>
            <div class="row mb-3">
                <label for="Status" class="col-sm-2 col-form-label">Status</label>
                <div class="col-sm-10">
                    <select name="status" class="form-control" required >
                        <option value="1">Aktif</option>
                        <option value="0">Non Aktif</option>
                    </select>
                </div>
            </div>
            <div class="row mb-3">                
                <div class="col-sm-10 offset-sm-2">
                    <button type="submit" class="btn btn-primary">Simpan</button>
                </div>
            </div>
        </form>
    </main>
        
</body>

</html>
{{end}}