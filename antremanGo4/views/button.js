const buttons = document.getElementsByClassName('saveButton');

for (let i = 0; i < buttons.length; i++) {
    buttons[i].addEventListener('click', function(event) {
        // Tıklanan butonun bulunduğu en yakın .content div'ini alıyoruz
        let parentDiv = event.target;
        
        // .closest() kullanmadan, parentElement ile en yakın .content div'ini buluyoruz
        while (parentDiv && !parentDiv.classList.contains('idea')) {
            parentDiv = parentDiv.parentElement;
        }

        if (parentDiv) {
           // Parent div içindeki verileri alıyoruz (querySelector kullanmadan) 
            const childDivs = parentDiv.getElementsByTagName('div');
            const childi = parentDiv.getElementsByTagName('i')
            const childi2 = childi[3];
            childi2.className = "fa-solid fa-bookmark"
            
            const nameDiv = childDivs[0];  // İlk div (name bilgisi)

            const idea = nameDiv.innerText;
            const data = {
                idea: idea
            };

            fetch('/save', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })
            .then(response => {
                if (response.ok) {
                    return response.text();
                } else {
                    throw new Error('Failed to save data');
                }
            })
            .then(data => {
                alert(data);
            })
            .catch(error => {
                console.error('Error:', error);
            });
        }
    });
}