const buttonss = document.getElementsByClassName('likebtn');

for (let i = 0; i < buttonss.length; i++) {
    buttonss[i].addEventListener('click', function(event) {
        // Tıklanan butonun bulunduğu en yakın .idea div'ini alıyoruz
        let parentDivv = event.target;

        // .closest() kullanmadan, parentElement ile en yakın .idea div'ini buluyoruz
        while (parentDivv && !parentDivv.classList.contains('idea')) {
            parentDivv = parentDivv.parentElement;
        }

        if (parentDivv) {
            const childDivss = parentDivv.getElementsByTagName('div');
            const childii3 = parentDivv.getElementsByTagName('span');
            const childii = parentDivv.getElementsByTagName('i');

            // İlk 'span' elemanını alıyoruz
            const childii4 = childii3[2];
            
            // İlk 'i' elemanını alıyoruz
            const childii2 = childii[0];

            if (childii2.classList.contains("fa-solid") && childii2.classList.contains("fa-heart")) {
                // 'i' elemanının sınıfını değiştiriyoruz
                childii2.className = "fa-regular fa-heart";

                // 'span' elemanının içeriğini alıyoruz ve sayıya dönüştürüyoruz
                let intVal = parseInt(childii4.textContent, 10); // veya childii4.innerText
                if (isNaN(intVal)) intVal = 0; // Eğer dönüşüm başarısızsa sıfırla

                // 'intVal' değerini azaltıyoruz
                intVal -= 1;

                // Güncellenmiş değeri 'span' elemanına geri yazıyoruz
                childii4.textContent = intVal;

                // parentDivv'nin data-id'sini al
                const ideaId = parentDivv.getAttribute('data-id');
                console.log(ideaId);

                const dataa = {
                    intval : intVal, // Anahtar adı 'Likenumberr' olarak ayarlanıyor
                    classname : childii2.className
                };

                // Template literals kullanarak URL'yi biçimlendiriyoruz
                fetch(`/likes/${ideaId}`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(dataa)
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

            } else {
                // 'i' elemanının sınıfını değiştiriyoruz
                childii2.className = "fa-solid fa-heart";
                childii2.style.color = "#e9071e";

                // 'span' elemanının içeriğini alıyoruz ve sayıya dönüştürüyoruz
                let intVal = parseInt(childii4.textContent, 10); // veya childii4.innerText
                if (isNaN(intVal)) intVal = 0; // Eğer dönüşüm başarısızsa sıfırla

                // 'intVal' değerini artırıyoruz
                intVal += 1;

                // Güncellenmiş değeri 'span' elemanına geri yazıyoruz
                childii4.textContent = intVal;

                // parentDivv'nin data-id'sini al
                const ideaId = parentDivv.getAttribute('data-id');
                console.log(ideaId);

                const dataa = {
                    intval : intVal, // Anahtar adı 'Likenumberr' olarak ayarlanıyor
                    classname : childii2.className
                };

                // Template literals kullanarak URL'yi biçimlendiriyoruz
                fetch(`/likes/${ideaId}`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(dataa)
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
        }
    });
}

