const ADODB = require('node-adodb');

// Connect to a database
const connection = ADODB.open('Provider=Microsoft.Jet.OLEDB.4.0;Data Source=database.mdb;');

(async () => {
    try {
        // Create a table
        await connection.execute('CREATE TABLE Users (ID COUNTER PRIMARY KEY, Name TEXT)');
        console.log('Table created.');

        // Insert data
        await connection.execute("INSERT INTO Users (Name) VALUES ('John Doe')");
        console.log('Data inserted.');

        // Query data
        const rows = await connection.query('SELECT * FROM Users');
        console.log('Data retrieved:', rows);
    } catch (error) {
        console.error(error);
    }
})();
