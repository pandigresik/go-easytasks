{{define "index"}}
<!DOCTYPE html>
<html>

<head>
    {{template "_header" .}}
</head>

<body>
    
    {{template "_menu" .}}
    
    
    <!-- Begin page content -->
    <main class="flex-shrink-0 p-4">
        <div>
            <button type="submit" class="btn btn-primary" onclick="window.location.href='/form'">Add task</button>
        </div>
        <table class="table table-bordered">
            <thead>
                <tr>
                    <th>Nama</th>
                    <th>Pic</th>
                    <th>Deadline</th>
                    <th>Status</th>
                    <th>Aksi</th>
                </tr>
            </thead>
            <tbody>
                {{ if .data }}
                {{ range .data }}
                <tr>                    
                    <td>{{ .Name }}</td>
                    <td>{{ .Pic }}</td>
                    <td>{{ .Deadline }}</td>
                    <td>{{ if .Status }} Aktif {{ else }} Non Aktif {{ end }}</td>
                    <td><a href="/editform/{{.ID }}" class="btn btn-primary" >Edit</a> <a href="/deleteTask/{{.ID }}" class="btn btn-danger" >Hapus</a></td>                    
                </tr>
                {{ end }}
                {{ end }}
            </tbody>
        </table>
    </main>
    
    </body>
</html>
{{end}}