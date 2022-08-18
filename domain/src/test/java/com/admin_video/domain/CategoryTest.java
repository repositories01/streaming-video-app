package com.admin_video.domain;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class CategoryTest {
    @Test
    public void TestNewCategory() {
        Assertions.assertNotNull(new Category());
    }
}
